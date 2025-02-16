package client

import (
	"log"
	"strconv"
	"strings"
	"time"

	"distivity/config/static"
	"distivity/types"
	"distivity/utils"

	"github.com/bwmarrin/discordgo"
)

var discordSession *discordgo.Session

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func colorize(color, message string) string {
	return color + message + colorReset
}

func setActivity(s *discordgo.Session, activity string, config types.Config) {
	if activity == "" {
		return
	}

	guild, err := s.State.Guild(config.Discord.GuildID)
	if err != nil {
		log.Printf(colorize(colorRed, "Error getting guild: %v"), err)
		return
	}

	activity = strings.Replace(activity, "{count}", strconv.Itoa(guild.MemberCount), 1)
	activity = utils.ReplaceEmojis(activity)

	err = s.UpdateCustomStatus(activity)
	if err != nil {
		log.Printf(colorize(colorRed, "Error setting activity: %v"), err)
	}
}

func presenceUpdateHandler(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	// Presence update detected
}

func InitDiscordClient() {
	config := static.GetConfig()
	log.Println(colorize(colorCyan, "Setting up the Sessions..."))
	var err error
	discordSession, err = discordgo.New("Bot " + config.Credentials.DiscordToken)
	if err != nil {
		log.Fatalf(colorize(colorRed, "Error creating Discord session: %v"), err)
	}

	discordSession.StateEnabled = true
	discordSession.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMembers)
	discordSession.Identify.Intents |= discordgo.IntentsGuildPresences
	discordSession.Identify.Intents |= discordgo.IntentsGuilds

	discordSession.AddHandler(presenceUpdateHandler)

	err = discordSession.Open()
	if err != nil {
		log.Fatalf(colorize(colorRed, "Error opening Discord session: %v"), err)
	}

	log.Println(colorize(colorGreen, "Discord session opened successfully"))

	err = discordSession.RequestGuildMembers(config.Discord.GuildID, "", 0, "", true)
	if err != nil {
		log.Printf(colorize(colorRed, "Error requesting guild members: %v"), err)
	} else {
		log.Println(colorize(colorGreen, "Guild members requested successfully"))
		log.Println(colorize(colorYellow, "----------------------------------------"))
	}

	log.Println(colorize(colorCyan, "Verifying state cache... (5 seconds)"))

	time.Sleep(5 * time.Second)

	log.Println("\033[H\033[2J")

	guild, err := discordSession.State.Guild(config.Discord.GuildID)
	if err != nil {
		log.Printf(colorize(colorRed, "Error verifying state cache: %v"), err)
	} else {
		log.Printf(colorize(colorGreen, "State cache verified: %d members"), guild.MemberCount)
	}

	setActivity(discordSession, config.Discord.CustomStatus, config)
}

func GetDiscordSession() *discordgo.Session {
	return discordSession
}
