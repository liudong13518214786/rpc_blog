package cache

import (
	"errors"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/cache"
	//_"github.com/astaxie/beego/cache/memcache"
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	server = "127.0.0.1:11211"
	mc     *memcache.Client
	Mcache *MemcacheHandler
)

//func Init()  {
//	mc, err:=cache.NewCache("memcache", "")
//}

type MemcacheHandler struct {
}

func init() {
	mc = memcache.New(server)
	if mc == nil {
		beego.Error("memcache connection failed")
		return
	}
	beego.Info("memcache connection success!")
}

func (m *MemcacheHandler) Set(key, value string, exp int32) bool {
	err := mc.Set(&memcache.Item{Key: key, Value: []byte(value), Expiration: exp})
	if err != nil {
		beego.Error("memcache set key error:", err)
		return false
	}
	return true
}

func (m *MemcacheHandler) Get(key string) ([]byte, error) {
	res, _ := mc.Get(key)
	if string(res.Key) == key {
		value := res.Value
		return value, nil
	}
	beego.Error("memcache get key error")
	err := errors.New("memcache get key error")
	return []byte(""), err
}

func (m *MemcacheHandler) Delete(key string) bool {
	err := mc.Delete(key)
	if err != nil {
		beego.Error(err)
		return false
	}
	return true
}
