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

	err = common.InitRedisDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Redis连接成功")

	err = common.InitMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Mysql连接成功")

	conf := config.NewServerConfig()

	server := http.Server{
		Addr:    conf.Address,
		Handler: conf.Handler,
	}

	controller.RegisterRoutes()

	log.Print("服务已启动...")
	server.ListenAndServe()
	log.Print("服务已关闭")

	common.CloseRedisDB()
	common.CloseMysqlDB()
}
