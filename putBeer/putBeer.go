package putbeer

import (
	db "github.com/AkapongAlone/komgrip-test/orm"
	"net/http"

	"github.com/gin-gonic/gin"
	
)

type result struct {
	Name  string
	Type string
	Picture	string
	Detail string
  }

func UpdateBeer(c *gin.Context) {
	var result result
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
	detailBeer.Type = result.Detail
	id := c.Param("id")

	var beforeEdit db.NameBeer
	db.Db.Where("id = ?", id).First(&beforeEdit)
	db.Db.Model(&beforeEdit).Updates(nameBeer)

	var beforeEditDetail db.DetailBeer
	db.Db.Where("id = ?", id).First(&beforeEditDetail)
	db.Db.Model(&beforeEditDetail).Updates(detailBeer)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}


