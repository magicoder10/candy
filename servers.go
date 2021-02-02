package main

import (
	"fmt"
	"net"

	"candy/observability"
	"candy/pubsub"
	"candy/server"
	"candy/server/proto"

	"google.golang.org/grpc"
)

func main() {
	logger := observability.NewLogger(observability.Info)

	pubSubServer := server.NewPubSub(&logger)
	pubSubServer.StartPubSub()

	pubSubRemote := pubsub.NewRemote(&logger)
	gameServer := server.NewGame(&logger, pubSubRemote)

	port := 8081

	go func() {
		server.WaitReady("localhost", port)
		pubSubRemote.Start("localhost", port)
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterPubSubServer(grpcServer, pubSubServer)
	proto.RegisterGameServer(grpcServer, &gameServer)

	logger.Infof("Server started at localhost:%d\n", port)

	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("Fail to start server: %w", err)
		panic(err)
	}
}
