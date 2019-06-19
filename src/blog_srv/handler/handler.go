package handler

import (
	"context"
	blog_ext "rpc_blog/proto/blog"
	"rpc_blog/src/blog_srv/db"
	"rpc_blog/src/module/utils"
)

type BlogServiceExtHandler struct {
}

func (b *BlogServiceExtHandler) GetBlogList(ctx context.Context, req *blog_ext.BlogListRequest, resp *blog_ext.BlogListResponse) error {
	userid := req.Userid
	_, status := utils.CheckParma(userid)
	if !status {
		resp.BlogList = nil
		return nil
	}
	bloglist, err := db.GetBlogList()
	if err != nil {
		resp.BlogList = nil
		return nil
	}
	resp.BlogList = bloglist
	return nil
}

func (b *BlogServiceExtHandler) WriteBlog(ctx context.Context, req *blog_ext.WriteBlogRequest, resp *blog_ext.WriteBlogResponse) error {
	UserId := req.Userid
	BlogId := req.BlogId
	res, success := utils.CheckParma(UserId)
	if !success {
		resp.Code = 500
		resp.Message = res
		return nil
	}
	if BlogId != "" {
		blog, err := db.GetBlogDetail(BlogId)
		if err != nil {
			resp.Code = 500
			resp.Message = "system error"
			return nil
		}
		if blog != nil {
			resp.Code = 505
			resp.Message = "blog id is exist"
			return nil
		}
	}
	err := db.WriteBlog(UserId, req.Title, req.Title)
	if err != nil {
		resp.Code = 510
		resp.Message = "write blog err"
		return nil
	}
	return nil
}

func (b *BlogServiceExtHandler) DeleteBlog(ctx context.Context, req *blog_ext.DeleteBlogRequest, resp *blog_ext.DeleteBlogResponse) error {
	bid := req.BlogId
	userid := req.Userid
	res, err := db.GetBlogDetailByUseruuid(bid, userid)
	if err != nil {
		resp.Code = 500
		resp.Message = "找不到bid"
	}
	return nil
}

func (b *BlogServiceExtHandler) ModifyBlog(ctx context.Context, req *blog_ext.ModifyBlogRequest, resp *blog_ext.ModifyBlogResponse) error {
	return nil
}

func (b *BlogServiceExtHandler) BlogDetail(ctx context.Context, req *blog_ext.BlogDetailRequest, resp *blog_ext.BlogDetailResponse) error {
	return nil
}
