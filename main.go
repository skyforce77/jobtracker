package main

import (
	"./providers"
	"./snapshot"
	"./view"
	"errors"
	"github.com/xconstruct/go-pushbullet"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"sort"
)

func main() {
	app := cli.NewApp()

	app.Name = "jobtracker"
	app.Usage = "Find your dream job"

	app.Commands = []cli.Command{
		{
			Name:      "snap",
			Usage:     "Save job offers in a file to use them later",
			ArgsUsage: "<snapshot> <providers>...",
			Action: func(c *cli.Context) error {
				if c.NArg() < 2 {
					return errors.New("wrong arguments")
				}

				snap := snapshot.NewSnapshot(c.Args()[0])
				pro := make([]providers.Provider, c.NArg()-1)

				for i := 1; i < c.NArg(); i++ {
					argProv := c.Args()[i]
					if argProv == "all" {
						pro = append(pro, providers.GetProviders()...)
						continue
					}

					provider := providers.ProviderFromName(argProv)
					if provider != nil {
						pro[i-1] = provider
					} else {
						return errors.New("provider not found")
					}
				}

				log.Println("Creating snapshot...")

				cn := make(chan error, 8)

				for _, p := range pro {
					sp := p
					go func() {
						err := sp.RetrieveJobs(snap.Collector())
						cn <- err
					}()
				}

				i := 0
				for i != len(pro) {
					err := <-cn
					log.Printf("Provider finished (%d/%d)\n", i+1, len(pro))
					if err != nil {
						log.Printf("Error detected: %s\n", err)
					}
					snap.Save()
					i++
				}
				log.Println("Snap", c.Args()[0], "created")
				return nil
			},
		},
		{
			Name:      "view",
			Usage:     "View available jobs",
			ArgsUsage: "[providers]...",
			Action: func(c *cli.Context) error {
				pro := make([]providers.Provider, 0)
				if c.IsSet("snapshot") {
					pro = append(pro, snapshot.NewSnapshot(c.String("snapshot")))
				}
				for _, name := range c.Args() {
					if name == "all" {
						pro = append(pro, providers.GetProviders()...)
						continue
					}

					provider := providers.ProviderFromName(name)
					if provider != nil {
						pro = append(pro, provider)
					}
				}
				view.StartView(pro)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "snapshot, s",
					Usage: "Load snapshot from `FILE`",
				},
			},
		},
		{
			Name:  "notify",
			Usage: "Notifies updates",
			Subcommands: []cli.Command{
				{
					Name:      "pushbullet",
					Usage:     "Notifies updates on pushbullet",
					ArgsUsage: "[pushbullet token] [snap name] [providers]...",
					Action: func(c *cli.Context) error {
						original := snapshot.NewSnapshot(c.Args()[1])
						snap := snapshot.NewSnapshot(c.Args()[1])
						snap.Erase()

						pro := make([]providers.Provider, 0)
						if c.IsSet("snapshot") {
							pro = append(pro, snapshot.NewSnapshot(c.String("snapshot")))
						}
						for _, name := range c.Args()[2:] {
							if name == "all" {
								pro = append(pro, providers.GetProviders()...)
								continue
							}

							provider := providers.ProviderFromName(name)
							if provider != nil {
								pro = append(pro, provider)
							}
						}

						log.Println("Creating snapshot...")

						cn := make(chan error, 8)

						for _, p := range pro {
							sp := p
							go func() {
								err := sp.RetrieveJobs(snap.Collector())
								cn <- err
							}()
						}

						i := 0
						for i != len(pro) {
							err := <-cn
							log.Printf("Provider finished (%d/%d)\n", i+1, len(pro))
							if err != nil {
								log.Printf("Error detected: %s\n", err)
							}
							snap.Save()
							i++
						}
						log.Println("Snap", c.Args()[1], "created")

						diff, err := providers.NewDiff(original, snap)
						if err != nil {
							return err
						}

						pb := pushbullet.New(c.Args()[0])
						devs, err := pb.Devices()
						if err != nil {
							panic(err)
						}

						for _, j := range diff.Added {
							for _, dev := range devs {
								dev.PushLink(j.Title+" - "+j.Company+" - "+j.Location, j.Link, j.Desc)
							}
						}

						snap.Save()
						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "snapshot, s",
							Usage: "Load snapshot from `FILE`",
						},
					},
				},
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
