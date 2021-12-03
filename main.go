package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"service-template/global"
	"service-template/internal/routers"
)

func init() {
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("setupSetting failed: %v", err)
	}

	err = global.DBConnectSetting()
	if err != nil {
		log.Fatalf("配置数据库连接失败：%v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	if global.ServerSetting.RunMode == "release" {
		gin.DisableConsoleColor()
	}
	router := routers.NewRouters()
	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Starting server failed: %v", err)
	}
}
