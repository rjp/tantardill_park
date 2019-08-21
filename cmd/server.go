package main

import (
	"bitbucket.org/rjp/ports/portrpc"
	"context"
	"fmt"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type ourClient struct{}

func (pdbs *ourClient) PutPort(ctx context.Context, port *portrpc.Port) (*portrpc.PutResponse, error) {
	fmt.Printf("%s\n", port.String())
	return &portrpc.PutResponse{Response: "yes"}, nil
}

func (pdbs *ourClient) GetPortByShortcode(ctx context.Context, shortcode *portrpc.Shortcode) (*portrpc.Port, error) {
	x := portrpc.Port{}
	return &x, nil
}

func main() {
	tcpport := 9387
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", tcpport))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	portrpc.RegisterPortDatabaseServer(grpcServer, &ourClient{})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
