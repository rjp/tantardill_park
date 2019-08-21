package main

import (
	"bitbucket.org/rjp/tantardill_park/portrpc"
	"context"
	"errors"
	"fmt"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

var Database map[string]portrpc.Port

type ourClient struct{}

func (pdbs *ourClient) PutPort(ctx context.Context, port *portrpc.Port) (*portrpc.PutResponse, error) {
	fmt.Printf("%s\n", port.String())
	Database[port.Shortcode] = *port
	return &portrpc.PutResponse{Response: "yes"}, nil

}

func (pdbs *ourClient) GetPortByShortcode(ctx context.Context, shortcode *portrpc.Shortcode) (*portrpc.Port, error) {
	fmt.Printf("HELLO SOMEONE WANTED %s\n", shortcode.Shortcode)
	if port, ok := Database[shortcode.Shortcode]; ok {
		return &port, nil
	} else {
		return &portrpc.Port{}, errors.New("No port for shortcode " + shortcode.Shortcode)
	}
}

func main() {
	// "III. Store config in the environment"
	// This also helps running it in Docker since we can easily
	// pass them when we're starting up a container.
	tcpport := os.Getenv("PORTS_GRPC_PORT")
	if tcpport == "" {
		// Panic or default? Default is nicer for testing.
		fmt.Println("Missing port, defaulting to 9387")
		tcpport = "9387"
	}

	Database = make(map[string]portrpc.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", tcpport))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	portrpc.RegisterPortDatabaseServer(grpcServer, &ourClient{})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
