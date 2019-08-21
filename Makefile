# This is probably an awful `Makefile` but I stopped using `make`
# in favour of `redo` about 10 years ago...
all: client server

portrpc/portrpc.pb.go: portrpc/portrpc.proto
	protoc  -I portrpc/ portrpc/portrpc.proto --go_out=plugins=grpc:portrpc

client: cmd/tantardill_park_client/main.go portrpc/portrpc.pb.go
	go build -o client cmd/tantardill_park_client/main.go

server: cmd/tantardill_park_server/main.go portrpc/portrpc.pb.go
	go build -o server cmd/tantardill_park_server/main.go

clean:
	rm -f client server portrpc/portrpc.pb.go
