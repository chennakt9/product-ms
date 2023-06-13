package main

import (
	"context"
	"log"
	"time"

	pb "github.com/chennakt9/go-grpc-setup/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("%v Response data", res)

}