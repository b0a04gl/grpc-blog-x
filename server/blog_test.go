package main

import (
    "context"
    "testing"
    "github.com/gRPCBlogX/api"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
)

type MockService struct {
    
}

func (s *MockService) CreateBlog(ctx context.Context, req *api.BlogRequest) (*api.BlogResponse, error) {
    
    return &api.BlogResponse{Id: "123"}, nil
}

func (s *MockService) GetBlog(ctx context.Context, req *api.GetBlogRequest) (*api.Blog, error) {
    
    if req.Id == "123" {
        return &api.Blog{Id: "123", Title: "Test Blog", Content: "Test Content"}, nil
    }
    return nil, grpc.Errorf(codes.NotFound, "Blog not found")
}

func (s *MockService) UpdateBlog(ctx context.Context, req *api.UpdateBlogRequest) (*api.BlogResponse, error) {
    
    if req.Id == "123" {
        return &api.BlogResponse{Id: "123"}, nil
    }
    return nil, grpc.Errorf(codes.NotFound, "Blog not found")
}

func (s *MockService) DeleteBlog(ctx context.Context, req *api.DeleteBlogRequest) (*api.DeleteBlogResponse, error) {
    
    if req.Id == "123" {
        return &api.DeleteBlogResponse{Success: true, Message: "Blog deleted"}, nil
    }
    return nil, grpc.Errorf(codes.NotFound, "Blog not found")
}

func (s *MockService) ListBlogs(ctx context.Context, req *api.ListBlogsRequest) (*api.ListBlogsResponse, error) {
    
    // Return a list of sample blogs
    blogs := []*api.Blog{
        {Id: "123", Title: "Test Blog 1", Content: "Test Content 1"},
        {Id: "124", Title: "Test Blog 2", Content: "Test Content 2"},
        {Id: "125", Title: "Test Blog 3", Content: "Test Content 3"},
    }
    
    return &api.ListBlogsResponse{Blogs: blogs}, nil
}

func TestCreateBlog(t *testing.T) {
    
    service := &MockService{}
    
    req := &api.BlogRequest{Title: "Test Blog", Content: "Test Content"}
    resp, err := service.CreateBlog(context.Background(), req)
    if err != nil {
        t.Errorf("CreateBlog failed: %v", err)
    }
    if resp.Id != "123" {
        t.Errorf("CreateBlog returned unexpected ID: %s", resp.Id)
    }
}

func TestGetBlog(t *testing.T) {
    
    service := &MockService{}
    
    req := &api.GetBlogRequest{Id: "123"}
    blog, err := service.GetBlog(context.Background(), req)
    if err != nil {
        t.Errorf("GetBlog failed: %v", err)
    }
    if blog.Id != "123" {
        t.Errorf("GetBlog returned unexpected ID: %s", blog.Id)
    }
}

func TestUpdateBlog(t *testing.T) {
    
    service := &MockService{}
    
    req := &api.UpdateBlogRequest{Id: "123", Title: "Updated Blog", Content: "Updated Content"}
    resp, err := service.UpdateBlog(context.Background(), req)
    if err != nil {
        t.Errorf("UpdateBlog failed: %v", err)
    }
    if resp.Id != "123" {
        t.Errorf("UpdateBlog returned unexpected ID: %s", resp.Id)
    }
}

func TestDeleteBlog(t *testing.T) {
    
    service := &MockService{}
    
    req := &api.DeleteBlogRequest{Id: "123"}
    resp, err := service.DeleteBlog(context.Background(), req)
    if err != nil {
        t.Errorf("DeleteBlog failed: %v", err)
    }
    if !resp.Success {
        t.Errorf("DeleteBlog returned unexpected success: false")
    }
}

func TestListBlogs(t *testing.T) {
    
    service := &MockService{}
    
    req := &api.ListBlogsRequest{}
    resp, err := service.ListBlogs(context.Background(), req)
    if err != nil {
        t.Errorf("ListBlogs failed: %v", err)
    }
    if len(resp.Blogs) != 3 {
        t.Errorf("ListBlogs returned unexpected number of blogs: %d", len(resp.Blogs))
    }
}
