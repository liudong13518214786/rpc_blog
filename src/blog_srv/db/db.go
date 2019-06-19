package db

import _ "github.com/lib/pq"
import "github.com/jmoiron/sqlx"

var (
	db *sqlx.DB
)

func InitDatabase(Dsn string) {
	db = sqlx.MustConnect("postgresql", Dsn)
	db.SetMaxIdleConns(1)  //设置最大空闲连接数
	db.SetMaxOpenConns(10) //设置最大的连接数
}
