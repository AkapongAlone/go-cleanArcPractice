package delbeer

import (
	"net/http"
	"strconv"

	db "github.com/AkapongAlone/komgrip-test/orm"

	"github.com/gin-gonic/gin"
)

func DelBeer(c *gin.Context) {
	idParam := c.Param("id")
	id,err :=strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
	}
	del := db.Db.Delete(&db.NameBeer{},id)
	if err := del.Error; err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	del = db.Db.Delete(&db.DetailBeer{},id-1)
	if err := del.Error; err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}
