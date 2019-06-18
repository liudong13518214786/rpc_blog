package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"rpc_blog/config"
	"strings"
	"time"
)

func GetTimeNowFormat() string {
	today := time.Now().Format("2006-01-02 15:04:05")
	return today
}

func CreateToken(email string) string {
	auth_raw := []byte(config.SECRETKEY + ":" + email)
	return base64.StdEncoding.EncodeToString(auth_raw)
}

func CheckToken(token string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		err := errors.New("base64 decode error")
		return "", err
	}
	result := string(res)
	resList := strings.Split(result, ":")
	if len(resList) != 2 {
		err := errors.New("token error 12")
		return "", err
	}
	key, email := resList[0], resList[1]
	if key != config.SECRETKEY {
		err := errors.New("token error 13")
		return "", err
	}
	return email, nil
}

func CheckParma(parm ...string) (string, bool) {
	//判断参数是否为空
	for _, v := range parm {
		if v == "" {
			res := fmt.Sprintf("参数%v为空", v)
			return res, false
		}
	}
	return "", true
}