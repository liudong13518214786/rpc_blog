package handler

import (
	"context"
	blog_ext "rpc_blog/proto/blog"
)

type BlogServiceExtHandler struct {
}

func (b *BlogServiceExtHandler) GetBlogList(ctx context.Context, req *blog_ext.BlogListRequest, resp *blog_ext.BlogListResponse) error {
	return nil
}

func (b *BlogServiceExtHandler) WriteBlog(ctx context.Context, req *blog_ext.WriteBlogRequest, resp *blog_ext.WriteBlogResponse) error {
	return nil
}

func (b *BlogServiceExtHandler) DeleteBlog(ctx context.Context, req *blog_ext.DeleteBlogRequest, resp *blog_ext.DeleteBlogResponse) error {
	return nil
}

func (b *BlogServiceExtHandler) ModifyBlog(ctx context.Context, req *blog_ext.ModifyBlogRequest, resp *blog_ext.ModifyBlogResponse) error {
	return nil
}

func (b *BlogServiceExtHandler) BlogDetail(ctx context.Context, req *blog_ext.BlogDetailRequest, resp *blog_ext.BlogDetailResponse) error {
	return nil
}