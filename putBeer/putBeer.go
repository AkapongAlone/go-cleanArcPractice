package putbeer

import (
	db "github.com/AkapongAlone/komgrip-test/orm"
	"net/http"

	"github.com/gin-gonic/gin"
	
)

func UpdateBeer(c *gin.Context) {
	var nameBeer db.NameBeer
	if err := c.ShouldBindJSON(&nameBeer); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := c.Param("id")

	var beforeEdit db.NameBeer
	db.Db.Where("id = ?", id).First(&beforeEdit)
	db.Db.Model(&beforeEdit).Updates(nameBeer)
	c.JSON(http.StatusOK, gin.H{
		"data": nameBeer,
	})
}

func UpdateDetail(c *gin.Context) {
	var detailBeer db.DetailBeer
	if err := c.ShouldBindJSON(&detailBeer); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := c.Param("id")
	var beforeEdit db.DetailBeer
	db.Db.Where("id = ?", id).First(&beforeEdit)
	db.Db.Model(&beforeEdit).Updates(detailBeer)
	c.JSON(http.StatusOK, gin.H{
		"data": detailBeer,
	})
}
