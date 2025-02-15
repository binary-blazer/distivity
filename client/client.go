package client

import (
	"log"
	"strconv"
	"strings"

	"distivity/config/static"
	"distivity/types"
	"distivity/utils"

	"github.com/bwmarrin/discordgo"
)

var discordSession *discordgo.Session

func setActivity(s *discordgo.Session, activity string, config types.Config) {
	if activity == "" {
		return
	}

	guild, err := s.State.Guild(config.Discord.GuildID)
	if err != nil {
		log.Printf("Error getting guild: %v", err)
		return
	}

	activity = strings.Replace(activity, "{count}", strconv.Itoa(guild.MemberCount), 1)
	activity = utils.ReplaceEmojis(activity)

	err = s.UpdateCustomStatus(activity)
	if err != nil {
		log.Printf("Error setting activity: %v", err)
	}
}

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

	setActivity(discordSession, config.Discord.CustomStatus, config)
}

func GetDiscordSession() *discordgo.Session {
	return discordSession
}

func activityHandler(s *discordgo.Session, m *discordgo.PresenceUpdate) {
	log.Printf("User %s is now %s", m.User.ID, m.Status)
}
