package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/gRPCBlogX/api"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	blogService := newBlogServer()

	pb.RegisterBlogServiceServer(server, blogService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Blog gRPC server is listening on port 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
