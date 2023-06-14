proto:
	protoc --go_out=. --go-grpc_out=. pkg/proto/greet.proto

server:
	go run cmd/main.go