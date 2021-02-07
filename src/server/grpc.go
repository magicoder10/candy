package server

import (
	"fmt"
	"net"

	"candy/observability"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	name       string
	logger     *observability.Logger
	grpcServer *grpc.Server
}

func (g GRPCServer) Start(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	g.logger.Infof("%s service started at localhost:%d\n", g.name, port)

	if err := g.grpcServer.Serve(lis); err != nil {
		g.logger.Fatalf("Fail to start %s service: %w", g.name, err)
		return err
	}
	return nil
}

func (g GRPCServer) Stop() {
	if g.grpcServer == nil {
		return
	}
	g.grpcServer.Stop()
}
