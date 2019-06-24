package db

import (
	"database/sql"
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
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return bloglist, err
}

func GetBlogDetail(bid string) (*Blog, error) {
	blog := &Blog{}
	err := db.Get(blog, "SELECT uuid, useruuid, info,title,build_time FROM blog WHERE uuid=$1 AND status=$2 ORDER BY build_time DESC ;", bid, "normal")
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return blog, err
}

func WriteBlog(useruuid, info, title string) error {
	buildTime := utils.GetTimeNowFormat()
	uuid := utils.GenerateRandomString("blg", 8)
	_, err := db.Exec("INSERT INTO blog(uuid, useruuid, info, title, build_time, status) VALUES ($1, $2, $3, $4, $5, $6);", uuid, useruuid, info, title, buildTime, "normal")
	return err
}

func GetBlogDetailByUseruuid(bid, useruuid string) (*Blog, error) {
	blog := &Blog{}
	err := db.Get(blog, "SELECT uuid, useruuid, info,title,build_time FROM blog WHERE uuid=$1 AND status=$2 AND useruuid=$3 ORDER BY build_time DESC ;", bid, "normal", useruuid)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return blog, err
}

func UpdateBlogStatus(bid, status string) error {
	_, err := db.Exec("UPDATE blog SET status=$1 WHERE uuid=$2 AND status=$3;", status, bid, "normal")
	return err
}

func UpdateBlogInfo(bid, title, info, useruuid string) error {
	_, err := db.Exec("UPDATE blog SET title=$1,info=$2 WHERE uuid=$3 AND status=$4 AND useruuid=$5;", title, info, bid, "normal", useruuid)
	return err
}
