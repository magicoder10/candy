package pubsub

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"io"

	"candy/observability"
	"candy/server/proto"

	"google.golang.org/grpc"
)

type Remote struct {
	logger  *observability.Logger
	conn    *grpc.ClientConn
	client  proto.PubSubClient
	context context.Context
}

func (r Remote) Subscribe(topic Topic, callback func(payload []byte)) error {
	req := proto.SubscribeRequest{Topic: string(topic)}
	subClient, err := r.client.Subscribe(r.context, &req)
	if err != nil {
		return err
	}

	go func() {
		for {
			res, err := subClient.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				r.logger.Errorln(err)
				return
			}

			callback(res.Payload)
		}
	}()
	return nil
}

func (r Remote) Publish(topic Topic, payload interface{}) error {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(payload)
	if err != nil {
		return err
	}

	req := proto.PublishRequest{
		Topic:   string(topic),
		Payload: buf.Bytes(),
	}
	_, err = r.client.Publish(r.context, &req)
	return err
}

func (r *Remote) Start(ipAddress string, port int) error {
	var err error

	r.conn, err = grpc.Dial(fmt.Sprintf("%s:%d", ipAddress, port), grpc.WithInsecure())
	if err != nil {
		return err
	}
	r.client = proto.NewPubSubClient(r.conn)
	r.logger.Infoln("Remote PubSub started")
	return nil
}

func (r Remote) Stop() {
	r.conn.Close()
}

func NewRemote(logger *observability.Logger) *Remote {
	return &Remote{
		logger:  logger,
		context: context.Background(),
	}
}
