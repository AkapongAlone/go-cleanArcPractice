package routes

import (
	"net/http"
	"os"

	db "github.com/AkapongAlone/komgrip-test/database"
	_ "github.com/AkapongAlone/komgrip-test/docs"
	"github.com/AkapongAlone/komgrip-test/middlewares"
	auth "github.com/AkapongAlone/komgrip-test/src/auth/repositories"
	_ "github.com/AkapongAlone/komgrip-test/src/beers/handler"
	"github.com/AkapongAlone/komgrip-test/src/beers/repositories"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	// add swagger
	
	handle := repositories.NewBeerHandler(db.Db)
	authHandle := auth.NewAuthHandler(db.Db)
	r.POST("/register",authHandle.Register)
	r.GET("/login",authHandle.Login)
	r.GET("/re_token",authHandle.RefreshToken)
	protected := r.Group("/auth", middlewares.Protect([]byte(os.Getenv("SECRET")))) 
	{
		beerGroup := protected.Group("/api/v1")
		beerGroup.GET("/beer", handle.GetBeer)
		beerGroup.GET("/beer/:id", handle.GetBeerByID)
		beerGroup.GET("/image/:id", handle.GetImage)
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		beerGroup.Use(middlewares.LoggingInfoMiddleware())
		{beerGroup.POST("/beer", handle.PostBeer)
		beerGroup.PUT("/beer/:id", handle.UpdateBeer)
		beerGroup.DELETE("/beer/:id", handle.DeleteBeer)
		}
	}
	
	r.StaticFS("/file/", http.Dir("images"))
	return r
}
