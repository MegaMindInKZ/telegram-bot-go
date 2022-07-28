package main

import (
	"log"
	"telegram-bot/clients/tgClient"
	"telegram-bot/config"
	"telegram-bot/storage/sqlite"

	event_consumer "telegram-bot/consumer/event-consumer"
	"telegram-bot/events/telegram"
)

const (
	batchSize = 100
)

func main() {

	configuration := config.MustLoad()

	storage := sqlite.New()
	eventsProcessor := telegram.New(
		tgClient.New(configuration.TgBotHost, configuration.TgBotToken),
		storage,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
	defer storage.Database.Close()
	// db := mongo.New(cfg.MongoConnectionString)
}
