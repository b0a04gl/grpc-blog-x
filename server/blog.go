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
