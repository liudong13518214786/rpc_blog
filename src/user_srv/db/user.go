package db

import "time"

type User struct {
	UserName string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

func InsertUser(username, email, password string) error {
	today := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec("INSERT INTO user(username,password,build_time,email) VALUES(?,?,?,?)", username, password, today, email)
	return err
}

func GetUserByEmail(email, password string) (*User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE email=? AND password=? LIMIT 1", email, password)
	return &user, err
}
