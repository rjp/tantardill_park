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
	"sort"
)

var Database map[string]portrpc.Port

type ourClient struct{}

// `PutPort` adds (or updates) a port object to our database.
func (pdbs *ourClient) PutPort(ctx context.Context, port *portrpc.Port) (*portrpc.PutResponse, error) {
	// TODO Put them into a real database, such as Postgres or even SQLite
	Database[port.Shortcode] = *port
	return &portrpc.PutResponse{Response: "ok"}, nil

}

// `GetShortcodes` retrieves a sorted list of all the known shortcodes
func (pdbs *ourClient) GetShortcodes(empty *portrpc.GetShortcodesRequest, stream portrpc.PortDatabase_GetShortcodesServer) error {
	// Collect our shortcodes into a `[]string` before sending in order
	// to sort them (which makes life easier.)
	codes := []string{}
	for code, _ := range Database {
		codes = append(codes, code)
	}

	sort.Strings(codes)

	// Now send out each code to the stream as a packed `Shortcode` RPC object.
	for _, code := range codes {
		shortcode := portrpc.Shortcode{Shortcode: code}
		if err := stream.Send(&shortcode); err != nil {
			return err
		}
	}

	// All done, success is in you.
	return nil
}

// `GetPortByShortcode` retrieves a single port by its shortcode (ie `ZARCB`)
// Call as `http://server:port/shortcode/ZARCB`
func (pdbs *ourClient) GetPortByShortcode(ctx context.Context, shortcode *portrpc.Shortcode) (*portrpc.Port, error) {
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
