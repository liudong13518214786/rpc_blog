package db

import (
	_ "github.com/lib/pq"
	"log"
)
import "github.com/jmoiron/sqlx"

var (
	db  *sqlx.DB
	err error
)

func InitDatabase(Dsn string) {
	db, err = sqlx.Connect("postgres", Dsn)
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(1)  //设置最大空闲连接数
	db.SetMaxOpenConns(10) //设置最大的连接数
}
