package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"rpc_blog/config"
	"strings"
	"time"
)

func GetTimeNowFormat() string {
	today := time.Now().Format("2006-01-02 15:04:05")
	return today
}

func CreateToken(useruuid string) string {
	auth_raw := []byte(config.SECRETKEY + ":" + useruuid)
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
	key, useruuid := resList[0], resList[1]
	if key != config.SECRETKEY {
		err := errors.New("token error 13")
		return "", err
	}
	return useruuid, nil
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

func GenerateRandomString(suff string, l int) string {
	bytes := "0123456789abcdefghijklmnopqrstuvwxyz"
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return fmt.Sprintf("%s_%s", suff, string(result))
}
