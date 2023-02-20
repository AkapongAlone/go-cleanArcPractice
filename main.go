package main

import (
	_ "fmt"

	db "github.com/AkapongAlone/komgrip-test/orm"
	"github.com/AkapongAlone/komgrip-test/getBeer"
	"github.com/AkapongAlone/komgrip-test/delBeer"
	"github.com/AkapongAlone/komgrip-test/postBeer"
	"github.com/AkapongAlone/komgrip-test/putBeer"
	"github.com/gin-gonic/gin"
	
)

func main() {

	db.InitDB()

	router := gin.Default()

	router.GET("/beer/:name/:limit",getBeer.GetBeer)
	router.POST("/beer",postbeer.PostBeer)
	router.POST("/beer_detail",postbeer.PostDetail)
	router.PUT("/beer/:id",putbeer.UpdateBeer)
	router.PUT("/beer_detail/:id",putbeer.UpdateDetail)
	router.DELETE("/beer/:id",delbeer.DelBeer)
	router.DELETE("/beer_detail/:id",delbeer.DelDetail)

	router.Run()
}