package main

import (
	"log"
	"net"

	"github.com/olanta/olanta/internal/server"
	"github.com/olanta/olanta/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server := server.NewServer()

	proto.RegisterScannerServiceServer(grpcServer, server)

	log.Println("Server is running on port 8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
