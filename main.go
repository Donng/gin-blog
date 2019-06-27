package main

import (
	"fmt"
	"gin-blog/models"
	"gin-blog/pkg/logging"
	"gin-blog/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"

	"gin-blog/pkg/setting"
)

func main() {
	// 初始化配置、数据库、日志模块
	setting.Setup()
	models.Setup()
	logging.Setup()

	endless.DefaultReadTimeOut = setting.Server.ReadTimeout
	endless.DefaultWriteTimeOut = setting.Server.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.Server.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}