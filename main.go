package main

import (
	_ "fmt"

	db "github.com/AkapongAlone/komgrip-test/database"
	"github.com/AkapongAlone/komgrip-test/routes"
	
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8080
// @BasePath /api
func main() {
	db.InitDB()
	router := routes.SetupRouter()
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*",}
	// router.Use(cors.New(config))
	router.Run()
}
