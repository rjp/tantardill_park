package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"bitbucket.org/rjp/ports/portrpc"
	grpc "google.golang.org/grpc"
)

// Import our JSON from a file
// TODO Pass the filename as a parameter.
func importJSON(client portrpc.PortDatabaseClient) {
	f, err := os.Open("ports.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	importJSONFromReader(f, client)
}

// Parse JSON from an `io.Reader` (actually an `os.File`) using the
// streaming decode method. This way we only ever have one `Port` model
// alive at a time and should never consume a large amount of RAM.
//
// This is extracted from `importJSON` because I did originally try handling
// multipart file upload as a method of import but that turned out to need
// the file writing to disk and I'm not sure the constraints allow that.
func importJSONFromReader(f *os.File, client portrpc.PortDatabaseClient) {
	dec := json.NewDecoder(f)

	// Skip past the opening `{`
	t, err := dec.Token()
	// Panic for now since if we can't parse the JSON, we're stuck anyway.
	if err != nil {
		panic(err)
	}

	for dec.More() {
		// Now, oddly, the key part of an object is actually a `Token`,
		// not a `string` as I originally thought. Which lead to much
		// confusion as I was trying to `dec.Decode(&stringvar)` here.
		key, err := dec.Token()
		if err != nil {
			panic(err)
		}

		// We get back an `interface` but we want the `string` part.
		shortcode := key.(string)

		// Back to normal style decoding for the object that's attached
		// to the key.
		var port portrpc.Port
		err = dec.Decode(&port)
		if err != nil {
			panic(err)
		}

		// Keep all the parts of the object together.
		port.Shortcode = shortcode

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		ret, err := client.PutPort(ctx, &port)
		if err != nil {
			panic(err)
		}

		// TODO Convert this to `log` or something similar.
		fmt.Printf("ret=%s key=%s port=%#v\n", ret.String(), t, port)
	}
}

func main() {
	// TODO Get this from flags
	serverAddr := "127.0.0.1:9387"

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		panic(err)
	}

	client := portrpc.NewPortDatabaseClient(conn)

	// Make this a lambda because we want to capture the `client` variable
	// without making it global. Not sure if there's a simple way to stick
	// `client` somewhere into the http handling code such that we can then
	// retrieve it in a handler.
	shortcodeHandler := func(w http.ResponseWriter, req *http.Request) {
		// We're using `/shortcode/ABCDE` as our URL and the simplest way
		// to extract our shortcode is to just chop off the front.
		shortcode := strings.Replace(req.URL.Path, "/shortcode/", "", 1)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		// Package up our shortcode into a gRPC appropriate message.
		shortcodeRPC := &portrpc.Shortcode{Shortcode: shortcode}

		port, err := client.GetPortByShortcode(ctx, shortcodeRPC)
		// This error catch is too broad. It'll give a 404 for any error.
		// TODO Clarify this to only return 404 when we can't find the port.
		if err != nil {
			http.Error(w, "Port not found for code "+shortcode, http.StatusNotFound)
			return
		}

		// We probably want the data back as JSON since we're a REST(ish) API.
		encodedPort, err := json.Marshal(feature)
		if err != nil {
			http.Error(w, "{\"error\":\"JSON marshalling failed\"}", http.StatusInternalServerError)
			return
		}

		// If we made it this far, we're good to ok with a 200 and the JSON encoding of the port.
		w.WriteHeader(http.StatusOK)
		w.Write(encodedPort)
	}

	// "First service(ClientAPI) should parse the JSON file"
	// TODO: Add a file upload import API call to allow updates without having
	// to restart the service (needs to store the uploaded file temporarily on
	// disk which might be a problem with the constraints.)
	importJSON(client)

	// TODO Add more API calls for different queries.
	http.HandleFunc("/shortcode/", shortcodeHandler)

	// And begin.
	_ = http.ListenAndServe(":8288", nil)
}
