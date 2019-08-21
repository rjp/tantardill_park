package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"bitbucket.org/rjp/tantardill_park/portrpc"
	grpc "google.golang.org/grpc"
)

// `importJSON` loads our port information from a fixed filename.
// TODO Pass the filename as a parameter.
func importJSON(client portrpc.PortDatabaseClient) {
	f, err := os.Open("/ports/ports.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	importJSONFromReader(f, client)
}

// `importJSONFromReader` parses JSON from an `os.File` using the
// streaming decode method. This way we only ever have one `Port` model
// alive at a time and should never consume a large amount of RAM.
//
// This is extracted from `importJSON` because I did originally try handling
// multipart file upload as a method of import but that turned out to need
// the file writing to disk and I'm not sure the constraints allow that.
// But since we now have reload functionality, we can add new ports anyway.
func importJSONFromReader(f *os.File, client portrpc.PortDatabaseClient) {
	dec := json.NewDecoder(f)

	// Skip past the opening `{`
	_, err := dec.Token()

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

		// We don't care what we get back from this RPC call as long as there's
		// no error.
		// TODO Maybe implement some kind of "updated/inserted/deleted" response.
		_, err = client.PutPort(ctx, &port)
		if err != nil {
			panic(err)
		}
	}
}

// `fetchCodes` gets the list of shortcodes from the RPC server and
// returns it as JSON.
func fetchCodes(client portrpc.PortDatabaseClient) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.GetShortcodes(ctx, &portrpc.GetShortcodesRequest{})
	if err != nil {
		return []string{}, err
	}

	codes := []string{}
	for {
		code, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return []string{}, err
		}
		codes = append(codes, code.Shortcode)
	}

	return codes, nil
}

func main() {
	// "III. Store config in the environment"
	// This also helps running it in Docker since we can easily
	// pass them when we're starting up a container. Although
	// we'd still need to expose them which makes life faffy.
	tcpport := os.Getenv("PORTS_GRPC_PORT")
	if tcpport == "" {
		// Panic or default? Default is nicer for testing.
		fmt.Println("Missing port, defaulting to 9387")
		tcpport = "9387"
	}
	// Same for the hostname - this simplifies things when we
	// spin both parts up using `docker-compose`.
	tcphost := os.Getenv("PORTS_GRPC_HOST")
	if tcphost == "" {
		tcphost = "127.0.0.1"
	}
	serverAddr := tcphost + ":" + tcpport

	// Not bothering with TLS here but a production server probably
	// wants TLS with client certs to be secure.
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(serverAddr, opts...)

	// Simpler to `panic` here than work out some complicated retry
	// scheme - let the supervisor handle that if we need it.
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
		encodedPort, err := json.Marshal(port)
		if err != nil {
			http.Error(w, "{\"error\":\"JSON marshalling failed\"}", http.StatusInternalServerError)
			return
		}

		// If we made it this far, we're good to ok with a 200 and the JSON encoding of the port.
		w.WriteHeader(http.StatusOK)
		w.Write(encodedPort)
	}

	// If we call `/reload/`, reload the ports in `/ports/ports.json` and submit
	// them to the database. This doesn't handle any missing ports (ie there's no
	// way to remove one currently.) (Again, a lambda to trap `client`.)
	reloadHandler := func(w http.ResponseWriter, req *http.Request) {
		importJSON(client)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\":\"OK\"}"))
	}

	// Lambda, etc. The error checking and JSON encoding from this and
	// `shortcodeHandler` could be extracted and combined into a helper.
	//
	codesHandler := func(w http.ResponseWriter, req *http.Request) {
		codes, err := fetchCodes(client)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"error\":\"NOTOK\"}"))
			return
		}

		encodedCodes, err := json.Marshal(codes)
		if err != nil {
			http.Error(w, "{\"error\":\"JSON marshalling failed\"}", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(encodedCodes)
	}

	// "First service(ClientAPI) should parse the JSON file"
	// TODO: Add a file upload import API call to allow updates without having
	// to restart the service (needs to store the uploaded file temporarily on
	// disk which might be a problem with the constraints.)
	// DONE: Or at least add an API call which forces a refresh of the JSON
	// if it's held on an external volume to the Docker service.
	importJSON(client)

	// DONE: Add more API calls for different queries.
	http.HandleFunc("/shortcode/", shortcodeHandler)
	http.HandleFunc("/reload/", reloadHandler)
	http.HandleFunc("/codes/", codesHandler)

	// And begin.
	_ = http.ListenAndServe(":8288", nil)
}
