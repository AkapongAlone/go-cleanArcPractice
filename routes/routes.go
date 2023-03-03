package routes

import (
	"net/http"

	db "github.com/AkapongAlone/komgrip-test/database"
	_ "github.com/AkapongAlone/komgrip-test/src/beers/handler"
	"github.com/AkapongAlone/komgrip-test/src/beers/repositories"
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})
	handle := repositories.NewBeerHandler(db.Db)
	r.GET("/beer",handle.GetBeer)
	r.POST("/beer",handle.PostBeer)
	r.PUT("/beer/:id",handle.UpdateBeer)
	r.DELETE("/beer/:id",handle.DeleteBeer)
	r.StaticFS("/file/",http.Dir("images"))
	return r
}