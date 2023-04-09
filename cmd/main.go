package main

import (
	"user_system/config"
	"user_system/internal/router"
)

func Init() {
	config.InitConfig()
}

func main() {
	Init()
	router.InitRouterAndServe()
}
