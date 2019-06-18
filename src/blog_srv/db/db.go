package db

import _ "github.com/mysql"
import "github.com/jmoiron/sqlx"

var (
	db *sqlx.DB
)

func InitDatabase(mysqlDsn string) {
	db = sqlx.MustConnect("mysql", mysqlDsn)
	db.SetMaxIdleConns(1)  //设置最大空闲连接数
	db.SetMaxOpenConns(10) //设置最大的连接数
}
