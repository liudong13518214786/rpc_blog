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
	_, err := db.GetBlogDetailByUseruuid(bid, userid)
	if err != nil {
		resp.Code = 500
		resp.Message = "找不到bid"
		return nil
	}
	err = db.UpdateBlogStatus(bid, "destroy")
	if err != nil {
		resp.Code = 505
		resp.Message = "删除失败"
		return nil
	}
	resp.Code = 100
	resp.Message = "删除成功"
	return nil
}

func (b *BlogServiceExtHandler) ModifyBlog(ctx context.Context, req *blog_ext.ModifyBlogRequest, resp *blog_ext.ModifyBlogResponse) error {
	userid := req.Userid
	bid := req.BlogId
	title := req.Title
	detail := req.Detail
	status, success := utils.CheckParma(userid, bid, title, detail)
	if !success {
		resp.Code = 500
		resp.Message = status
		return nil
	}
	err := db.UpdateBlogInfo(bid, title, detail, userid)
	if err != nil {
		resp.Code = 505
		resp.Message = "更新失败"
		return nil
	}
	resp.Code = 100
	resp.Message = "修改成功"
	return nil
}

func (b *BlogServiceExtHandler) BlogDetail(ctx context.Context, req *blog_ext.BlogDetailRequest, resp *blog_ext.BlogDetailResponse) error {
	bid := req.BlogId
	//userid:=req.Userid
	_, success := utils.CheckParma(bid)
	if !success {
		resp.Detail = ""
		resp.Title = ""
		resp.Uuid = ""
		resp.BuildTime = ""
		return nil
	}
	blog, err := db.GetBlogDetail(bid)
	if err != nil {
		resp.Detail = ""
		resp.Title = ""
		resp.Uuid = ""
		resp.BuildTime = ""
		return nil
	}
	resp.Detail = blog.Info
	resp.Title = blog.Title
	resp.Uuid = blog.Uuid
	resp.BuildTime = blog.BuildTime
	return nil

}
