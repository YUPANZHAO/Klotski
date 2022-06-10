package main

import (
	"KlotskiWeb/common"
	"KlotskiWeb/controller"
	"KlotskiWeb/middleware"
	"log"
	"net/http"
)

func main() {
	var err error
	//启动redis
	err = common.InitRedisDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Redis连接成功")
	//启动mysql
	err = common.InitMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Mysql连接成功")
	//配置服务器
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &middleware.AuthMiddleware{},
	}
	//加载静态资源
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("wwwroot"))))
	//注册路由
	controller.RegisterRoutes()
	//启动监听
	log.Print("服务已启动...")
	server.ListenAndServe()
	log.Print("服务已关闭")
	//关闭连接
	common.CloseRedisDB()
}
