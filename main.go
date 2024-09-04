package main

import (
	"fmt"
	"k8s-plantform/config"
	"k8s-plantform/controller"
	// "k8s-plantform/db"
	"k8s-plantform/middle"
	"k8s-plantform/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// // 初始化数据库
	// db.Init()
	// 初始化k8s client
	service.K8s.Init()
	// 初始化gin
	r := gin.Default()
	// 加载jwt中间件
	//r.Use(middle.JWTAuth())
	// 加载跨域中间件
	r.Use(middle.Cors())
	// 初始化路由
	controller.Router.InitApiRouter(r)
	//启动websocket
	go func() {
		http.HandleFunc("/ws", service.Terminal.WsHandler)
		http.ListenAndServe(":8081", nil)
	}()
	// gin 程序启动
	fmt.Println("http://192.168.31.1:9091/")
	r.Run(config.ListenAddr)
	// // 关闭数据库
	// db.Close()
}
