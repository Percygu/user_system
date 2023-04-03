package main

import (
	log "github.com/sirupsen/logrus"
	"user_system/config"
	"user_system/internal/router"
)

func Init() {
	config.InitConfig()
}

func main() {
	Init()
	log.Info("111111111111")
	router.InitRouterAndServe()
}
