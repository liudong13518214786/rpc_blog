package db

import "rpc_blog/src/module/utils"

type User struct {
	UserId   string `json:"uuid" db:"uuid"`
	UserName string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

func InsertUser(username, email, password string) error {
	today := utils.GetTimeNowFormat()
	userid := utils.GenerateRandomString("usr", 8)
	_, err := db.Exec("INSERT INTO user(uuid, username,password,build_time,email) VALUES(?,?,?,?,?)", userid, username, password, today, email)
	return err
}

func GetUserByEmailPassword(email, password string) (*User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE email=? AND password=? LIMIT 1", email, password)
	return &user, err
}

func GetUserByEmail(email string) (*User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE email=? LIMIT 1", email)
	return &user, err
}

func GetUserByUuid(uuid string) (*User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE uuid=? LIMIT 1", uuid)
	return &user, err
}
