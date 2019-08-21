package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"bitbucket.org/rjp/ports/portrpc"
	grpc "google.golang.org/grpc"
)

func importJSON(client portrpc.PortDatabaseClient) {
	f, err := os.Open("ports.json")
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(f)

	t, err := dec.Token()
	if err != nil {
		panic(err)
	}
	fmt.Printf("TOKEN %#v\n", t)

	for dec.More() {
		key, err := dec.Token()
		if err != nil {
			panic(err)
		}
		fmt.Printf("TOKEN %#v\n", key)
		t := key.(string)

		var port portrpc.Port
		err = dec.Decode(&port)
		if err != nil {
			panic(err)
		}

		port.Shortcode = t

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		ret, err := client.PutPort(ctx, &port)
		if err != nil {
			panic(err)
		}

		fmt.Printf("ret=%s key=%s port=%#v\n", ret.String(), t, port)
	}
}

func main() {
	serverAddr := "127.0.0.1:9387"
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := portrpc.NewPortDatabaseClient(conn)

	importJSON(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	shortcode := &portrpc.Shortcode{Shortcode: "1234"}
	feature, err := client.GetPortByShortcode(ctx, shortcode)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", feature)

}
