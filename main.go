package main

import (
	"KlotskiWeb/common"
	"KlotskiWeb/config"
	"KlotskiWeb/controller"
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
	//获取服务器配置
	conf := config.NewServerConfig()
	//配置服务器
	server := http.Server{
		Addr:    conf.Address,
		Handler: conf.Handler,
	}
	//注册路由
	controller.RegisterRoutes()
	//启动监听
	log.Print("服务已启动...")
	server.ListenAndServe()
	log.Print("服务已关闭")
	//关闭连接
	common.CloseRedisDB()
	common.CloseMysqlDB()
}
