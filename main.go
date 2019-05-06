package main

import (
	"errors"
	"github.com/gregdel/pushover"
	"log"
	"os"
	"sort"

	"./discord"
	"./providers"
	"./snapshot"
	"./view"
	"github.com/xconstruct/go-pushbullet"
	"gopkg.in/urfave/cli.v1"
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
			Action:    actionSnap,
		},
		{
			Name:      "view",
			Usage:     "View available jobs",
			ArgsUsage: "[providers]...",
			Action:    actionView,
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
					Action:    actionPushBullet,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "snapshot, s",
							Usage: "Load snapshot from `FILE`",
						},
					},
				},
				{
					Name:      "pushover",
					Usage:     "Notifies updates on pushover",
					ArgsUsage: "[api token] [user token] [snap name] [providers]...",
					Action:    actionPushOver,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "snapshot, s",
							Usage: "Load snapshot from `FILE`",
						},
					},
				},
				{
					Name:      "discord",
					Usage:     "Notifies updates on discord",
					ArgsUsage: "[discord token] [discord channel_id] [snap name] [providers]...",
					Action:    actionDiscord,
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

func actionDiscord(c *cli.Context) error {
	original := snapshot.NewSnapshot(c.Args()[2])
	snap := snapshot.NewSnapshot(c.Args()[2])
	snap.Erase()

	pro := make([]providers.Provider, 0)
	if c.IsSet("snapshot") {
		pro = append(pro, snapshot.NewSnapshot(c.String("snapshot")))
	}
	for _, name := range c.Args()[3:] {
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
	snap.CollectFrom(pro, func(i int, err error) {
		log.Printf("Provider finished (%d/%d)\n", i+1, len(pro))
		if err != nil {
			log.Printf("Error detected: %s\n", err)
		}
		snap.Save()
	})

	log.Println("Snap", c.Args()[2], "created")
	diff, err := providers.NewDiff(original, snap)
	if err != nil {
		return err
	}
	err = discord.SendAndForget(c.Args()[0], c.Args()[1], diff, snap)
	if err != nil {
		return err
	}
	return nil
}

func actionPushBullet(c *cli.Context) error {
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
	snap.CollectFrom(pro, func(i int, err error) {
		log.Printf("Provider finished (%d/%d)\n", i+1, len(pro))
		if err != nil {
			log.Printf("Error detected: %s\n", err)
		}
		snap.Save()
	})

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

	snap.Save()

	for _, j := range diff.Added {
		for _, dev := range devs {
			err := dev.PushLink("Job offer alert",
				j.Link,
				j.Title+" - "+j.Company+" - "+j.Location)
			if err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}

func actionPushOver(c *cli.Context) error {
	original := snapshot.NewSnapshot(c.Args()[2])
	snap := snapshot.NewSnapshot(c.Args()[2])
	snap.Erase()

	pro := make([]providers.Provider, 0)
	if c.IsSet("snapshot") {
		pro = append(pro, snapshot.NewSnapshot(c.String("snapshot")))
	}
	for _, name := range c.Args()[3:] {
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
	snap.CollectFrom(pro, func(i int, err error) {
		log.Printf("Provider finished (%d/%d)\n", i+1, len(pro))
		if err != nil {
			log.Printf("Error detected: %s\n", err)
		}
		snap.Save()
	})

	log.Println("Snap", c.Args()[2], "created")

	diff, err := providers.NewDiff(original, snap)
	if err != nil {
		return err
	}

	app := pushover.New(c.Args()[0])
	recipient := pushover.NewRecipient(c.Args()[1])

	snap.Save()

	for _, j := range diff.Added {
		message := pushover.NewMessageWithTitle(j.Title+" - "+j.Company+" - "+j.Location,
			"Job offer alert")
		message.URL = j.Link

		_, err := app.SendMessage(message, recipient)
		if err != nil {
			log.Panic(err)
		}
	}

	return nil
}

func actionView(c *cli.Context) error {
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
}

func actionSnap(c *cli.Context) error {
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
}
