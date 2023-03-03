package main

import (
	_ "fmt"

	db "github.com/AkapongAlone/komgrip-test/database"
	"github.com/AkapongAlone/komgrip-test/routes"
)

func main() {
	db.InitDB()
	router := routes.SetupRouter()
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*",}
	// router.Use(cors.New(config))
	router.Run()
}
