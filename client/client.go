package client

import (
	"log"

	"distivity/config/static"

	"github.com/bwmarrin/discordgo"
)

var discordSession *discordgo.Session

func InitDiscordClient() {
	config := static.GetConfig()
	var err error
	discordSession, err = discordgo.New("Bot " + config.Credentials.DiscordToken)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	discordSession.AddHandler(activityHandler)

	err = discordSession.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}
}

func GetDiscordSession() *discordgo.Session {
	return discordSession
}

func activityHandler(s *discordgo.Session, m *discordgo.PresenceUpdate) {
	log.Printf("User %s is now %s", m.User.ID, m.Status)
}
