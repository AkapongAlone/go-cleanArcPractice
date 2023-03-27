package middlewares

import (
	"context"
	_ "errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	db "github.com/AkapongAlone/komgrip-test/database"
	"github.com/AkapongAlone/komgrip-test/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Ownerror struct {
	Code	int `json:"code"`
	Message string `json:"message"`
}
func LoggingInfoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().Format(time.RFC1123)
		c.Next()

		collection := db.GetMONGO()
		logInfo := &models.LogInfo{
			RequestAt:  start,
			StatusCode: c.Writer.Status(),
			Path:       c.Request.RequestURI,
			Method:     c.Request.Method,
			ClientAddr: c.ClientIP(),
		}

		_, err := collection.InsertOne(context.Background(), logInfo)
		if err != nil {
			log.Fatalln(err.Error())
		}

	}
}

func Protect(signature []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenString string
		if auth := ctx.Request.Header.Get("Authorization"); auth != "" {
			tokenString = strings.TrimPrefix(auth, "Bearer ")
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return

		}
		// refreshToken(&tokenString)
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing")
			}
			return signature, nil
		})
		var newErr Ownerror 
		if err != nil {
			newErr.Code = 401
			newErr.Message = "error from json"
			ctx.JSON(401,newErr)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		if claims["type"] != "access" {
			fmt.Println("Invalid token type")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return 
		}
		ctx.Next()
	}
}
