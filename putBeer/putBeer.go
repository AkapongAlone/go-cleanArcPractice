package putbeer

import (
	"net/http"
	"strconv"

	db "github.com/AkapongAlone/komgrip-test/orm"

	"github.com/gin-gonic/gin"
)


func UpdateBeer(c *gin.Context) {
	var result db.Result
	var nameBeer db.NameBeer
	var detailBeer db.DetailBeer

	if err := c.ShouldBindJSON(&result); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	nameBeer.Name = result.Name
	nameBeer.Type = result.Type
	nameBeer.Picture = result.Picture
	detailBeer.Type = result.Type
	detailBeer.Detail = result.Detail
	idParam := c.Param("id")
	id,err :=strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
	}
	var beforeEdit db.NameBeer
	db.Db.Where("id = ?", id).First(&beforeEdit)
	db.Db.Model(&beforeEdit).Updates(nameBeer)

	var beforeEditDetail db.DetailBeer
	db.Db.Where("id = ?", id-1).First(&beforeEditDetail)
	db.Db.Model(&beforeEditDetail).Updates(detailBeer)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}


