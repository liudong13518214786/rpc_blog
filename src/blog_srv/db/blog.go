package db

type Blog struct {
	Uuid      string `json:"uuid" db:"uuid"`
	UserUuid  string `json:"useruuid" db:"useruuid"`
	Info      string `json:"info" db:"info"`
	Title     string `json:"title" db:"title"`
	BuildTime string `json:"build_time" db:"build_time"`
}

func GetBlogList() ([]Blog, error) {
	var bloglist []Blog
	err := db.Select(bloglist, "SELECT uuid, useruuid, info,title,build_time FROM blog WHERE status=? ORDER BY build_time DESC ;", "normal")
	return bloglist, err
}

func GetBlogDetail(bid string) (*Blog, error) {
	blog := &Blog{}
	err := db.Get(blog, "SELECT uuid, useruuid, info,title,build_time FROM blog WHERE uuid=? AND status=? ORDER BY build_time DESC ;", bid, "normal")
	return blog, err
}
