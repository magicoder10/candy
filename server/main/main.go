package main

import (
	"fmt"
	"net"
	"time"

	"candy/observability"
	"candy/pubsub"
	"candy/server"
)

func main() {
	logger := observability.NewLogger(observability.Info)

	pubSubRemote := pubsub.NewRemote(&logger)
	pubSubServer := server.NewPubSub(&logger)

	gameServer := server.NewGame(&logger, pubSubRemote)

	go func() {
		err := pubSubServer.Start(8081)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		waitReady("localhost", 8081)
		pubSubRemote.Start("localhost", 8081)
	}()

	err := gameServer.Start(8082)
	if err != nil {
		panic(err)
	}
}

func waitReady(ipAddress string, port int) {
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ipAddress, port), time.Second)
		if err, ok := err.(*net.OpError); ok && err.Timeout() {
			conn.Close()
			continue
		}
		conn.Close()
		return
	}
}
