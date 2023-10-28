package main

import (
    "context"
    "fmt"
    "log"
    "google.golang.org/grpc"

    pb "github.com/gRPCBlogX/api"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewBlogServiceClient(conn)

    // Call gRPC methods
    resp, err := client.CreateBlog(context.TODO(), &pb.BlogRequest{
        Title:   "My First Blog",
        Content: "This is the content of my first blog.",
    })
    if err != nil {
        log.Fatalf("Error creating blog: %v", err)
    }
    fmt.Printf("Blog created with ID: %s\n", resp.Id)

    getResp, err := client.GetBlog(context.TODO(), &pb.GetBlogRequest{
        Id: resp.Id,
    })
    if err != nil {
        log.Fatalf("Error getting blog: %v", err)
    }
    fmt.Printf("Blog retrieved: %v\n", getResp)
}
