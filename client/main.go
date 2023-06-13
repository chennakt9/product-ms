package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/chennakt9/go-grpc-setup/proto"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost" + port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)


	callSayHello(client)

}