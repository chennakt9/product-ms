package main

import (
	"log"
	"net"

	"github.com/chennakt9/product-ms/pkg/db"
	pb "github.com/chennakt9/product-ms/pkg/pb"
	services "github.com/chennakt9/product-ms/pkg/services"
	"google.golang.org/grpc"
)

const (
	port = "0.0.0.0:50052"
)

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start the server, %v", err)
	}

	h := db.Init("host=productdb.stage.svc.cluster.local user=chenna password=1234 dbname=productdb port=5432")

	s := services.Server {
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	log.Printf("Server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}