package db

import (
	blog_ext "rpc_blog/proto/blog"
	"rpc_blog/src/module/utils"
)

type Blog struct {
	Uuid      string `json:"uuid" db:"uuid"`
	UserUuid  string `json:"useruuid" db:"useruuid"`
	Info      string `json:"info" db:"info"`
	Title     string `json:"title" db:"title"`
	BuildTime string `json:"build_time" db:"build_time"`
}

func GetBlogList() ([]*blog_ext.BlogInfo, error) {
	var bloglist []*blog_ext.BlogInfo
	err := db.Select(&bloglist, "SELECT uuid,title,build_time as buildTime FROM blog WHERE status='normal' ORDER BY build_time DESC ;")
	return bloglist, err
}

func GetBlogDetail(bid string) (*Blog, error) {
	blog := &Blog{}
	err := db.Get(blog, "SELECT uuid, useruuid, info,title,build_time FROM blog WHERE uuid=? AND status=? ORDER BY build_time DESC ;", bid, "normal")
	return blog, err
}

func WriteBlog(useruuid, info, title string) error {
	buildTime := utils.GetTimeNowFormat()
	uuid := utils.GenerateRandomString("blg", 8)
	_, err := db.Exec("INSERT INTO blog(uuid, useruuid, info, title, build_time, status) VALUES (?, ?, ?, ?, ?, ?);", uuid, useruuid, info, title, buildTime, "normal")
	return err
}

func GetBlogDetailByUseruuid(bid, useruuid string) (*Blog, error) {
	blog := &Blog{}
	err := db.Get(blog, "SELECT uuid, useruuid, info,title,build_time FROM blog WHERE uuid=? AND status=? AND useruuid=? ORDER BY build_time DESC ;", bid, "normal", useruuid)
	return blog, err
}

func UpdateBlogStatus(bid, status string) error {
	_, err := db.Exec("UPDATE blog SET status=? WHERE bid=? AND status=?;", status, bid, "normal")
	return err
}

func UpdateBlogInfo(bid, title, info, useruuid string) error {
	_, err := db.Exec("UPDATE blog SET title=?,info=? WHERE bid=? AND status=? AND useruuid=?;", title, info, bid, "normal", useruuid)
	return err
}
