package main

import (
	"context"
	"fmt"

	pb "github.com/gRPCBlogX/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type blogServer struct {
	blogs map[string]*pb.Blog
	pb.UnimplementedBlogServiceServer
}

func newBlogServer() *blogServer {
	return &blogServer{
		blogs: make(map[string]*pb.Blog),
	}
}

func (s *blogServer) CreateBlog(ctx context.Context, req *pb.BlogRequest) (*pb.BlogResponse, error) {

	id := fmt.Sprint(len(s.blogs) + 1)

	blog := &pb.Blog{
		Id:      id,
		Title:   req.Title,
		Content: req.Content,
	}
	s.blogs[id] = blog
	return &pb.BlogResponse{Id: id}, nil
}

func (s *blogServer) GetBlog(ctx context.Context, req *pb.GetBlogRequest) (*pb.Blog, error) {
	blog, exists := s.blogs[req.Id]
	if !exists {
		return nil, grpc.Errorf(codes.NotFound, "Blog not found")
	}
	return blog, nil
}

func (s *blogServer) UpdateBlog(ctx context.Context, req *pb.UpdateBlogRequest) (*pb.BlogResponse, error) {
	blog, exists := s.blogs[req.Id]
	if !exists {
		return nil, grpc.Errorf(codes.NotFound, "Blog not found")
	}

	blog.Title = req.Title
	blog.Content = req.Content

	return &pb.BlogResponse{Id: blog.Id}, nil
}

func (s *blogServer) DeleteBlog(ctx context.Context, req *pb.DeleteBlogRequest) (*pb.DeleteBlogResponse, error) {
	if _, exists := s.blogs[req.Id]; !exists {
		return nil, grpc.Errorf(codes.NotFound, "Blog not found")
	}

	delete(s.blogs, req.Id)

	return &pb.DeleteBlogResponse{Success: true, Message: "Blog deleted successfully"}, nil
}

func (s *blogServer) ListBlogs(ctx context.Context, req *pb.ListBlogsRequest) (*pb.ListBlogsResponse, error) {
	blogs := make([]*pb.Blog, 0)
	for _, blog := range s.blogs {
		blogs = append(blogs, blog)
	}

	return &pb.ListBlogsResponse{Blogs: blogs}, nil
}
