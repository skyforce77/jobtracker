package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/skyforce77/jobtracker/providers"
	"github.com/skyforce77/jobtracker/snapshot"
)

// SendAndForget creates a discord client and send new jobs offer
func SendAndForget(token string, channelID string, diff *providers.Diff, snap *snapshot.Snapshot) error {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Println("Error while creating Discord session: ", err)
		return err
	}

	sc := make(chan error, 1)
	dg.AddHandler(func(s *discordgo.Session, event *discordgo.Ready) {
		snap.Save()
		for _, j := range diff.Added {
			_, err := s.ChannelMessageSend(
				channelID,
				j.Title+" - "+j.Company+" - "+j.Location+" - "+j.Link,
			)
			if err != nil {
				log.Println("Error while creating Discord session: ", err)
				sc <- err
			}
		}
		dg.Close()
		sc <- nil
	})

	err = dg.Open()
	if err != nil {
		fmt.Println("Error while opening connection: ", err)
		return err
	}

	err = <-sc
	if err != nil {
		fmt.Println("Error while sending messages: ", err)
		return err
	}
	return nil
}
