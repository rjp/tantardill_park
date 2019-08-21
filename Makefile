# This is probably an awful `Makefile` but I stopped using `make`
# in favour of `redo` about 10 years ago...
all: client server

portrpc/portrpc.pb.go: portrpc/portrpc.proto
	protoc  -I portrpc/ portrpc/portrpc.proto --go_out=plugins=grpc:portrpc

client: cmd/client.go portrpc/portrpc.pb.go
	go build -o client cmd/client.go

server: cmd/server.go portrpc/portrpc.pb.go
	go build -o client cmd/server.go

clean:
	rm -f client server portrpc/portrpc.pb.go
