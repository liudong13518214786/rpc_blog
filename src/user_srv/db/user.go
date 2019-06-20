package db

import (
	"database/sql"
	"rpc_blog/src/module/utils"
)

type User struct {
	UserId   string `json:"uuid" db:"uuid"`
	UserName string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

func InsertUser(username, email, password string) error {
	today := utils.GetTimeNowFormat()
	userid := utils.GenerateRandomString("usr", 8)
	_, err := db.Exec("INSERT INTO users(uuid, username,password,build_time,email,status) VALUES($1,$2,$3,$4,$5,$6)", userid, username, password, today, email, "normal")
	return err
}

func GetUserByEmailPassword(email, password string) (*User, error) {
	var user User
	err := db.Get(&user, "SELECT uuid, username, password, email FROM users WHERE email=$1 AND password=$2 LIMIT 1", email, password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.Get(&user, "SELECT uuid,username,password,email FROM users WHERE email=$1 LIMIT 1", email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func GetUserByUuid(uuid string) (*User, error) {
	var user User
	err := db.Get(&user, "SELECT uuid, username, password, email FROM users WHERE uuid=$1 LIMIT 1", uuid)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}
