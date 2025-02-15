package utils

import "strings"

func ReplaceEmojis(activity string) string {
	emojiMap := map[string]string{
		":smile:":      "😄",
		":heart:":      "❤️",
		":thumbsup:":   "👍",
		":thumbsdown:": "👎",
		":ok:":         "👌",
		":clap:":       "👏",
		":pray:":       "🙏",
		":100:":        "💯",
		":fire:":       "🔥",
		":tada:":       "🎉",
		":rocket:":     "🚀",
		":grin:":       "😁",
		":joy:":        "😂",
		":sob:":        "😭",
		":wink:":       "😉",
		":sunglasses:": "😎",
		":thinking:":   "🤔",
		":sleeping:":   "😴",
		":party:":      "🥳",
		":confused:":   "😕",
		":angry:":      "😠",
		":star:":       "⭐",
		":moon:":       "🌙",
		":sun:":        "☀️",
		":rainbow:":    "🌈",
		":unicorn:":    "🦄",
	}

	for code, emoji := range emojiMap {
		activity = strings.ReplaceAll(activity, code, emoji)
	}

	return activity
}
