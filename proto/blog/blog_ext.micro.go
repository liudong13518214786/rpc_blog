// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: go/src/rpc_blog/proto/blog/blog_ext.proto

package blog_ext

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for BlogService service

type BlogService interface {
	//查看博客，博客列表
	GetBlogList(ctx context.Context, in *BlogListRequest, opts ...client.CallOption) (*BlogListResponse, error)
	//写博客
	WriteBlog(ctx context.Context, in *WriteBlogRequest, opts ...client.CallOption) (*WriteBlogResponse, error)
	//删除博客
	DeleteBlog(ctx context.Context, in *DeleteBlogRequest, opts ...client.CallOption) (*DeleteBlogResponse, error)
	//修改博客
	ModifyBlog(ctx context.Context, in *ModifyBlogRequest, opts ...client.CallOption) (*ModifyBlogResponse, error)
	//博客详情
	BlogDetail(ctx context.Context, in *BlogDetailRequest, opts ...client.CallOption) (*BlogDetailResponse, error)
}

type blogService struct {
	c    client.Client
	name string
}

func NewBlogService(name string, c client.Client) BlogService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "blogservice"
	}
	return &blogService{
		c:    c,
		name: name,
	}
}

func (c *blogService) GetBlogList(ctx context.Context, in *BlogListRequest, opts ...client.CallOption) (*BlogListResponse, error) {
	req := c.c.NewRequest(c.name, "BlogService.GetBlogList", in)
	out := new(BlogListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogService) WriteBlog(ctx context.Context, in *WriteBlogRequest, opts ...client.CallOption) (*WriteBlogResponse, error) {
	req := c.c.NewRequest(c.name, "BlogService.WriteBlog", in)
	out := new(WriteBlogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogService) DeleteBlog(ctx context.Context, in *DeleteBlogRequest, opts ...client.CallOption) (*DeleteBlogResponse, error) {
	req := c.c.NewRequest(c.name, "BlogService.DeleteBlog", in)
	out := new(DeleteBlogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogService) ModifyBlog(ctx context.Context, in *ModifyBlogRequest, opts ...client.CallOption) (*ModifyBlogResponse, error) {
	req := c.c.NewRequest(c.name, "BlogService.ModifyBlog", in)
	out := new(ModifyBlogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogService) BlogDetail(ctx context.Context, in *BlogDetailRequest, opts ...client.CallOption) (*BlogDetailResponse, error) {
	req := c.c.NewRequest(c.name, "BlogService.BlogDetail", in)
	out := new(BlogDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BlogService service

type BlogServiceHandler interface {
	//查看博客，博客列表
	GetBlogList(context.Context, *BlogListRequest, *BlogListResponse) error
	//写博客
	WriteBlog(context.Context, *WriteBlogRequest, *WriteBlogResponse) error
	//删除博客
	DeleteBlog(context.Context, *DeleteBlogRequest, *DeleteBlogResponse) error
	//修改博客
	ModifyBlog(context.Context, *ModifyBlogRequest, *ModifyBlogResponse) error
	//博客详情
	BlogDetail(context.Context, *BlogDetailRequest, *BlogDetailResponse) error
}

func RegisterBlogServiceHandler(s server.Server, hdlr BlogServiceHandler, opts ...server.HandlerOption) error {
	type blogService interface {
		GetBlogList(ctx context.Context, in *BlogListRequest, out *BlogListResponse) error
		WriteBlog(ctx context.Context, in *WriteBlogRequest, out *WriteBlogResponse) error
		DeleteBlog(ctx context.Context, in *DeleteBlogRequest, out *DeleteBlogResponse) error
		ModifyBlog(ctx context.Context, in *ModifyBlogRequest, out *ModifyBlogResponse) error
		BlogDetail(ctx context.Context, in *BlogDetailRequest, out *BlogDetailResponse) error
	}
	type BlogService struct {
		blogService
	}
	h := &blogServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&BlogService{h}, opts...))
}

type blogServiceHandler struct {
	BlogServiceHandler
}

func (h *blogServiceHandler) GetBlogList(ctx context.Context, in *BlogListRequest, out *BlogListResponse) error {
	return h.BlogServiceHandler.GetBlogList(ctx, in, out)
}

func (h *blogServiceHandler) WriteBlog(ctx context.Context, in *WriteBlogRequest, out *WriteBlogResponse) error {
	return h.BlogServiceHandler.WriteBlog(ctx, in, out)
}

func (h *blogServiceHandler) DeleteBlog(ctx context.Context, in *DeleteBlogRequest, out *DeleteBlogResponse) error {
	return h.BlogServiceHandler.DeleteBlog(ctx, in, out)
}

func (h *blogServiceHandler) ModifyBlog(ctx context.Context, in *ModifyBlogRequest, out *ModifyBlogResponse) error {
	return h.BlogServiceHandler.ModifyBlog(ctx, in, out)
}

func (h *blogServiceHandler) BlogDetail(ctx context.Context, in *BlogDetailRequest, out *BlogDetailResponse) error {
	return h.BlogServiceHandler.BlogDetail(ctx, in, out)
}
