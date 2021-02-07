package main

import (
	"candy/env"
	"candy/observability"
	"candy/pubsub"
	"candy/server"
)

func main() {
	type pubSubConfig struct {
		PubSubServerHost string `env:"PUBSUB_SERVER_HOST" default:"localhost"`
		PubSubServerPort int    `env:"PUBSUB_SERVER_PORT" default:"8000"`
	}

	config := pubSubConfig{}
	err := env.ParseConfigFromEnv(&config)
	if err != nil {
		panic(err)
	}

	logger := observability.NewLogger(observability.Info)

	pubSubRemote := pubsub.NewRemote(&logger)

	gameServer := server.NewGame(&logger, pubSubRemote)

	go func() {
		server.WaitReady(config.PubSubServerHost, config.PubSubServerPort)
		pubSubRemote.Start(config.PubSubServerHost, config.PubSubServerPort)
	}()

	err = gameServer.Start(8082)
	if err != nil {
		panic(err)
	}
}
