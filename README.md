# A mini-micro-service cluster

(The name, `tantardill_park` was suggested by a
["Neural Network Placename Generator"](https://colinmorris.github.io/rbm/geo/))

## Summary

Starting from no gRPC experience, with years old Docker knowledge, and only the
vaguest whiff of Go modules, this was the hardest technical test I've done
(and one of the previous was writing a spreadsheet in Perl with cell functions.)

There's little testing of the business logic because there's basically nothing
useful on the internet about how to test gRPC code without spinning everything
up. However, I have added a small testing script which uses `docker-compose`
to spin everything up and run through a few simple tests of the basic functionality.

The Docker files aren't optimised due to my code layout - because both the client
and server depend on `portrpc/portrpc.pb.go`, they both tend to get rebuilt if
anything changes. It might have been better to split client and server into
entirely distinct packages with their own repos and move `portrpc` into another.

I'm using "shortcode" to mean "the 5 letter code the JSON file is keyed on".

## Running

```
docker-compose up --build
```

This gives you the client on `http://localhost:8288` loaded with `ports.json`.

If you prefer, you can build the parts outside of Docker using the `Makefile`.
That'll produce `client` and `server` binaries which do the exact same as they
would if they were inside Docker containers.

Both the client and server use the environment variable `PORTS_GRPC_PORT` to
set their gRPC port number (defaults to 9387.) The client additionally uses
`PORTS_GRPC_HOST` for the gRPC hostname of the server (defaults to `localhost`.)

## Server API

## PutPort

Adds a port to the database, keyed on the shortcode.

## GetPortByShortcode

Retrieves a single port matching the given shortcode.

## GetShortcodes

Gets all stored shortcodes.

## Client API calls

### Reloading

If you've changed `ports.json`, you can ask the client to refresh the server
with the new information.

```
curl http://localhost:8288/reload/
```

Returns a simple message to indicate success.
```
{"status":"OK"}
```

Currently it doesn't handle deletion but the best option there is probably
to add a call to the server for a full reset and then reinsertion rather
than trying to implicate synchronisation logic. Or add a simple `DeletePort`
call to the server, have the client call `GetShortcodes` before loading
the JSON, and then have it delete anything not loaded after parsing is done.

## Getting a single port via shortcode

```
curl http://localhost:8288/shortcode/ZARCB
```

Returns the information for, e.g., Richards Bay.
```
{"name":"Richards Bay","city":"Richards Bay","country":"South Africa","coordinates":[32.0382856,-28.7807276],
 "province":"KwaZulu-Natal","timezone":"Africa/Johannesburg","unlocs":["ZARCB"],"shortcode":"ZARCB"}
```

If the shortcode isn't found, a 404 response is returned.

## Getting all the shortcodes

```
curl http://localhost:8288/codes/
```

Returns a sorted list of all the shortcodes.
```
["AEAJM","AEAUH","AEDXB", ... many entries ... ,"ZWBUQ","ZWHRE","ZWUTA"]
```
