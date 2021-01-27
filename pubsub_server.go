package main

import (
	"candy/observability"
	"candy/server"
)

func main() {
	logger := observability.NewLogger(observability.Info)
	pubSubServer := server.NewPubSub(&logger)

	err := pubSubServer.Start(8081)
	if err != nil {
		panic(err)
	}
}
