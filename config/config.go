package config

import (
	"flag"
	"log"
)

type Config struct {
	TgBotToken string
	TgBotHost  string
}

func MustLoad() Config {
	tgBotToken := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	tgBotHost := flag.String(
		"tg-bot-host",
		"",
		"host for connection to telegram bot",
	)
	flag.Parse()

	if *tgBotToken == "" {
		log.Fatal("token is not specified")
	}
	return Config{
		TgBotToken: *tgBotToken,
		TgBotHost:  *tgBotHost,
	}
}
