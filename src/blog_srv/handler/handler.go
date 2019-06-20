package handler

import (
	"context"
	"errors"
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
		err := errors.New("缺少参数userid")
		return err
	}
	bloglist, err := db.GetBlogList()
	if err != nil {
		resp.BlogList = nil
		return err
	}
	resp.BlogList = bloglist
	return nil
}

func (b *BlogServiceExtHandler) WriteBlog(ctx context.Context, req *blog_ext.WriteBlogRequest, resp *blog_ext.WriteBlogResponse) error {
	UserId := req.Userid
	BlogId := req.BlogId
	title := req.Title
	detail := req.Info
	res, success := utils.CheckParma(UserId, title, detail)
	if !success {
		resp.Code = 500
		resp.Message = res
		err := errors.New(res)
		return err
	}
	if BlogId != "" {
		blog, err := db.GetBlogDetail(BlogId)
		if err != nil {
			resp.Code = 500
			resp.Message = "system error"
			err := errors.New("system error")
			return err
		}
		if blog != nil {
			resp.Code = 505
			resp.Message = "blog id is exist"
			err := errors.New("blog id is exist")
			return err
		}
	}
	err := db.WriteBlog(UserId, detail, title)
	if err != nil {
		resp.Code = 510
		resp.Message = "write blog err"
		return err
	}
	resp.Code = 100
	resp.Message = "write blog success"
	return nil
}

func (b *BlogServiceExtHandler) DeleteBlog(ctx context.Context, req *blog_ext.DeleteBlogRequest, resp *blog_ext.DeleteBlogResponse) error {
	bid := req.BlogId
	userid := req.Userid
	_, err := db.GetBlogDetailByUseruuid(bid, userid)
	if err != nil {
		resp.Code = 500
		resp.Message = "找不到bid"
		return err
	}
	err = db.UpdateBlogStatus(bid, "destroy")
	if err != nil {
		resp.Code = 505
		resp.Message = "删除失败"
		return err
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
		err := errors.New(status)
		return err
	}
	err := db.UpdateBlogInfo(bid, title, detail, userid)
	if err != nil {
		resp.Code = 505
		resp.Message = "更新失败"
		return err
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
		err := errors.New("缺少bid")
		return err
	}
	blog, err := db.GetBlogDetail(bid)
	if err != nil {
		resp.Detail = ""
		resp.Title = ""
		resp.Uuid = ""
		resp.BuildTime = ""
		err := errors.New("找不到内容")
		return err
	}
	resp.Detail = blog.Info
	resp.Title = blog.Title
	resp.Uuid = blog.Uuid
	resp.BuildTime = blog.BuildTime
	return nil
}
