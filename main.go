package main

import (
	"./providers"
	"./snapshot"
	"./view"
	"errors"
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

					provider := providers.ProviderFromName(argProv)
					if provider != nil {
						pro[i-1] = provider
					} else {
						return errors.New("provider not found")
					}
				}

				for _, p := range pro {
					p.RetrieveJobs(snap.Collector())
				}
				snap.Save()
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
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
