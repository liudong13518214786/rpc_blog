//package main
//
//import (
//	"github.com/astaxie/beego"
//	"net/http"
//)
//
//func handleRPC(w http.ResponseWriter, r *http.Request)  {
//	if r.URL.Path == "/"{
//		_, _=w.Write([]byte("OK, this is server"))
//		return
//	}
//	//todo 跨域处理
//	if r.Method == "OPTIONS"{
//		return
//	}
//	HandlerRpcJSOn(w, r)
//	return
//}
//
//func HandlerRpcJSOn(w http.ResponseWriter, r *http.Request)  {
//
//}
//
////程序入口
//func main()  {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", handleRPC)
//	_ = http.ListenAndServe(":8082", mux)
//	beego.Info("Listen on :8082")
//}
//

package main

import (
	"github.com/astaxie/beego"
	_ "rpc_blog/src/api_srv/routers"
)

func main() {
	beego.Run("127.0.0.1:8877")
}
