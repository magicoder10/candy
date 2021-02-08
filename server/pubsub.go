package server

import (
	"context"

	"candy/observability"
	"candy/pubsub"
	"candy/server/proto"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
)

var _ proto.PubSubServer = (*PubSub)(nil)

type PubSub struct {
	GRPCServer
	pubSub *pubsub.PubSub
}

func (p PubSub) Subscribe(request *proto.SubscribeRequest, server proto.PubSub_SubscribeServer) error {
	payloads := make(chan interface{})

	sub, err := p.pubSub.Subscribe(pubsub.Topic(request.Topic), func(payload interface{}) {
		payloads <- payload
	})
	if err != nil {
		return err
	}

	p.logger.Infof("Subscribed:%s\n", request.Topic)

	for payload := range payloads {
		p.logger.Infof("Received message:%s\n", request.Topic)

		res := proto.SubscribeResponse{
			Payload: payload.([]byte),
		}
		err = server.Send(&res)
		if err != nil {
			continue
		}
	}

	sub.Unsubscribe()
	return nil
}

func (p PubSub) Publish(ctx context.Context, request *proto.PublishRequest) (*empty.Empty, error) {
	p.pubSub.Publish(pubsub.Topic(request.Topic), request.Payload)
	p.logger.Infof("Published message:%s\n", request.Topic)
	return &empty.Empty{}, nil
}

func (p *PubSub) Start(port int) error {
	p.grpcServer = grpc.NewServer()
	proto.RegisterPubSubServer(p.grpcServer, p)

	p.pubSub.Start()

	p.GRPCServer.Start(port)
	return nil
}

func NewPubSub(logger *observability.Logger) PubSub {
	return PubSub{
		GRPCServer: GRPCServer{
			name:   "PubSub",
			logger: logger,
		},

		pubSub: pubsub.NewPubSub(logger),
	}
}
