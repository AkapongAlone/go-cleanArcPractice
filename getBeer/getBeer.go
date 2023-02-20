package getBeer

import (
	"net/http"
	"strconv"

	db "github.com/AkapongAlone/komgrip-test/orm"
	"github.com/gin-gonic/gin"
)

type result struct {
	Name  string
	Type string
	Picture	string
	Detail string
  }

func GetBeer(c *gin.Context) {
	
	var result []result
	nameParam := c.Query("name")
    limitParam := c.Query("limit")
	limit,err :=strconv.Atoi(limitParam)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	

	

	db.Db.Table("name_beers").Select("name_beers.name, name_beers.type, name_beers.picture, detail_beers.detail").
	Joins("left join detail_beers on detail_beers.type = name_beers.type").
	Where("name_beers.name = ?", nameParam).Limit(limit).Scan(&result)

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})



}