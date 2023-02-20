package main

import (
	_ "fmt"

	"github.com/AkapongAlone/komgrip-test/delBeer"
	"github.com/AkapongAlone/komgrip-test/getBeer"
	db "github.com/AkapongAlone/komgrip-test/orm"
	"github.com/AkapongAlone/komgrip-test/postBeer"
	"github.com/AkapongAlone/komgrip-test/putBeer"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*",}
	router.Use(cors.New(config))

	router.GET("/beer",getBeer.GetBeer)
	router.POST("/beer",postbeer.PostBeer)
	router.PUT("/beer/:id",putbeer.UpdateBeer)
	router.DELETE("/beer/:id",delbeer.DelBeer)

	router.Run()
}