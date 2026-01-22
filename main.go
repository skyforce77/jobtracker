package main

import (
	"errors"
	"log"
	"os"
	"reflect"

	"github.com/gregdel/pushover"
	"github.com/skyforce77/jobtracker/discord"
	"github.com/skyforce77/jobtracker/print"
	"github.com/skyforce77/jobtracker/providers"
	"github.com/skyforce77/jobtracker/snapshot"
	"github.com/skyforce77/jobtracker/util"
	"github.com/urfave/cli/v2"
	"github.com/xconstruct/go-pushbullet"
)

func main() {
	app := &cli.App{
		Name:  "jobtracker",
		Usage: "Find your dream job",
		Commands: []*cli.Command{
			{
				Name:      "snap",
				Usage:     "Save job offers in a file to use them later",
				ArgsUsage: "<snapshot> <providers>...",
				Action:    actionSnap,
			},
			{
				Name:      "print",
				Usage:     "Print available jobs",
				ArgsUsage: "[providers]...",
				Action:    actionPrint,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "snapshot",
						Aliases: []string{"s"},
						Usage:   "Load snapshot from `FILE`",
					},
					&cli.BoolFlag{
						Name:    "pretty",
						Aliases: []string{"p"},
						Value:   true,
						Usage:   "Prints in a pretty way",
					},
					&cli.StringFlag{
						Name:    "filter",
						Aliases: []string{"f"},
						Usage:   "Filters from `LUA_FILE`",
					},
				},
			},
			{
				Name:  "notify",
				Usage: "Notifies updates",
				Subcommands: []*cli.Command{
					{
						Name:      "pushbullet",
						Usage:     "Notifies updates on pushbullet",
						ArgsUsage: "[pushbullet token] [snap name] [providers]...",
						Action:    actionPushBullet,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "snapshot",
								Aliases: []string{"s"},
								Usage:   "Load snapshot from `FILE`",
							},
							&cli.StringFlag{
								Name:    "filter",
								Aliases: []string{"f"},
								Usage:   "Filters from `LUA_FILE`",
							},
						},
					},
					{
						Name:      "pushover",
						Usage:     "Notifies updates on pushover",
						ArgsUsage: "[api token] [user token] [snap name] [providers]...",
						Action:    actionPushOver,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "snapshot",
								Aliases: []string{"s"},
								Usage:   "Load snapshot from `FILE`",
							},
							&cli.StringFlag{
								Name:    "filter",
								Aliases: []string{"f"},
								Usage:   "Filters from `LUA_FILE`",
							},
						},
					},
					{
						Name:      "discord",
						Usage:     "Notifies updates on discord",
						ArgsUsage: "[discord token] [discord channel_id] [snap name] [providers]...",
						Action:    actionDiscord,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "snapshot",
								Aliases: []string{"s"},
								Usage:   "Load snapshot from `FILE`",
							},
							&cli.StringFlag{
								Name:    "filter",
								Aliases: []string{"f"},
								Usage:   "Filters from `LUA_FILE`",
							},
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func actionDiscord(c *cli.Context) error {
	original := snapshot.NewSnapshot(c.Args().Get(2))
	snap := snapshot.NewSnapshot(c.Args().Get(2))
	snap.Erase()

	pro := make([]providers.Provider, 0)
	if c.IsSet("snapshot") {
		pro = append(pro, snapshot.NewSnapshot(c.String("snapshot")))
	}
	for _, name := range c.Args().Slice()[3:] {
		if name == "all" {
			pro = append(pro, providers.GetProviders()...)
			continue
		}

		provider := providers.ProviderFromName(name)
		if provider != nil {
			pro = append(pro, provider)
		}
	}

	collector := snap.Collector()
	if c.IsSet("filter") {
		collector = util.Filter(c.String("filter"), collector)
	}

	log.Println("Creating snapshot...")
	snap.CollectFrom(pro, collector, func(i int, err *snapshot.ProviderFailure) {
		log.Printf("Provider finished (%d/%d)\n", i+1, len(pro))
		if err.Error != nil {
			log.Printf("Error detected %s: %s\n", reflect.TypeOf(err.Provider), err.Error)
		}
		snap.Save()
	})

	log.Println("Snap", c.Args().Get(2), "created")
	diff, err := providers.NewDiff(original, snap)
	if err != nil {
		return err
	}
	err = discord.SendAndForget(c.Args().Get(0), c.Args().Get(1), diff, snap)
	if err != nil {
		return err
	}
	return nil
}

func actionPushBullet(c *cli.Context) error {
	original := snapshot.NewSnapshot(c.Args().Get(1))
	snap := snapshot.NewSnapshot(c.Args().Get(1))
	snap.Erase()

	pro := make([]providers.Provider, 0)
	if c.IsSet("snapshot") {
		pro = append(pro, snapshot.NewSnapshot(c.String("snapshot")))
	}
	for _, name := range c.Args().Slice()[2:] {
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

	collector := snap.Collector()
	if c.IsSet("filter") {
		collector = util.Filter(c.String("filter"), collector)
	}

	snap.CollectFrom(pro, collector, func(i int, err *snapshot.ProviderFailure) {
		log.Printf("Provider finished (%d/%d)\n", i+1, len(pro))
		if err.Error != nil {
			log.Printf("Error detected %s: %s\n", reflect.TypeOf(*err.Provider).String(), err.Error)
		}
		snap.Save()
	})

	log.Println("Snap", c.Args().Get(1), "created")

	diff, err := providers.NewDiff(original, snap)
	if err != nil {
		return err
	}

	pb := pushbullet.New(c.Args().Get(0))
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
	original := snapshot.NewSnapshot(c.Args().Get(2))
	snap := snapshot.NewSnapshot(c.Args().Get(2))
	snap.Erase()

	pro := make([]providers.Provider, 0)
	if c.IsSet("snapshot") {
		pro = append(pro, snapshot.NewSnapshot(c.String("snapshot")))
	}
	for _, name := range c.Args().Slice()[3:] {
		if name == "all" {
			pro = append(pro, providers.GetProviders()...)
			continue
		}

		provider := providers.ProviderFromName(name)
		if provider != nil {
			pro = append(pro, provider)
		}
	}

	collector := snap.Collector()
	if c.IsSet("filter") {
		collector = util.Filter(c.String("filter"), collector)
	}

	log.Println("Creating snapshot...")
	snap.CollectFrom(pro, collector, func(i int, err *snapshot.ProviderFailure) {
		log.Printf("Provider finished (%d/%d)\n", i+1, len(pro))
		if err.Error != nil {
			log.Printf("Error detected %s: %s\n", reflect.TypeOf(*err.Provider).String(), err.Error)
		}
		snap.Save()
	})

	log.Println("Snap", c.Args().Get(2), "created")

	diff, err := providers.NewDiff(original, snap)
	if err != nil {
		return err
	}

	app := pushover.New(c.Args().Get(0))
	recipient := pushover.NewRecipient(c.Args().Get(1))

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

func actionPrint(c *cli.Context) error {
	pro := make([]providers.Provider, 0)
	if c.IsSet("snapshot") {
		pro = append(pro, snapshot.NewSnapshot(c.String("snapshot")))
	}
	for _, name := range c.Args().Slice() {
		if name == "all" {
			pro = append(pro, providers.GetProviders()...)
			continue
		}

		provider := providers.ProviderFromName(name)
		if provider != nil {
			pro = append(pro, provider)
		}
	}

	var filter *string = nil
	if c.IsSet("filter") {
		f := c.String("filter")
		filter = &f
	}

	pretty := c.Bool("pretty")
	print.Print(pro, pretty, filter)
	return nil
}

func actionSnap(c *cli.Context) error {
	if c.NArg() < 2 {
		return errors.New("wrong arguments")
	}

	snap := snapshot.NewSnapshot(c.Args().Get(0))
	pro := make([]providers.Provider, c.NArg()-1)

	for i := 1; i < c.NArg(); i++ {
		argProv := c.Args().Get(i)
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

	cn := make(chan *snapshot.ProviderFailure, 8)

	for _, p := range pro {
		sp := p
		go func() {
			if sp == nil {
				return
			}
			err := &snapshot.ProviderFailure{Provider: &sp, Error: sp.RetrieveJobs(snap.Collector())}
			cn <- err
		}()
	}

	i := 0
	for i != len(pro) {
		err := <-cn
		log.Printf("Provider finished (%d/%d)\n", i+1, len(pro))
		if err.Error != nil {
			log.Printf("Error detected %s: %s\n", reflect.TypeOf(*err.Provider).String(), err.Error)
		}
		errn := snap.Save()
		if errn != nil {
			log.Printf("Error detected : %s\n", errn)
		}
		i++
	}
	log.Println("Snap", c.Args().Get(0), "created")
	return nil
}
