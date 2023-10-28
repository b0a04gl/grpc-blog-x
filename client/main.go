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

	// Create 
	createResp1, err1 := client.CreateBlog(context.TODO(), &pb.BlogRequest{
		Title:   "My First Blog",
		Content: "This is the content of my first blog.",
	})
	createResp2, err2 := client.CreateBlog(context.TODO(), &pb.BlogRequest{
		Title:   "My Second Blog",
		Content: "This is the content of my second blog.",
	})
	if err1 != nil {
		log.Fatalf("Error creating blog: %v", err1)
	}
	if err2 != nil {
		log.Fatalf("Error creating blog: %v", err2)
	}

	fmt.Printf("Blog created with ID: %s\n", createResp1.Id)
	fmt.Printf("Blog created with ID: %s\n", createResp2.Id)
	// Get 
	getResp, err := client.GetBlog(context.TODO(), &pb.GetBlogRequest{
		Id: createResp1.Id,
	})
	if err != nil {
		log.Fatalf("Error getting blog: %v", err)
	}
	fmt.Printf("Blog retrieved: %v\n", getResp)

	// Update 
	updateResp, err := client.UpdateBlog(context.TODO(), &pb.UpdateBlogRequest{
		Id:      createResp2.Id,
		Title:   "Updated Blog Title",
		Content: "This is the updated content of the blog.",
	})
	if err != nil {
		log.Fatalf("Error updating blog: %v", err)
	}
	fmt.Printf("Blog updated with ID: %s\n", updateResp.Id)

	// Delete 
	deleteResp, err := client.DeleteBlog(context.TODO(), &pb.DeleteBlogRequest{
		Id: createResp1.Id,
	})
	if err != nil {
		log.Fatalf("Error deleting blog: %v", err)
	}
	if deleteResp.Success {
		fmt.Printf("Blog deleted with id : %v successfully\n",createResp1.Id)
	} else {
		fmt.Printf("Failed to delete blog: %s\n", deleteResp.Message)
	}

	// List
	listResp, err := client.ListBlogs(context.TODO(), &pb.ListBlogsRequest{})
	if err != nil {
		log.Fatalf("Error listing blogs: %v", err)
	}
	fmt.Printf("List of Blogs:\n")
	for _, blog := range listResp.Blogs {
		fmt.Printf("ID: %s, Title: %s, Content: %s\n", blog.Id, blog.Title, blog.Content)
	}
}
