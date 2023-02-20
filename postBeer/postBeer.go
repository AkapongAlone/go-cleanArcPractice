package postbeer

import (
	_"fmt"
	"net/http"
	"strconv"
	"time"

	db "github.com/AkapongAlone/komgrip-test/orm"
	"github.com/gin-gonic/gin"
)



func PostBeer(c *gin.Context) {
	var result db.Result
	var nameBeer db.NameBeer
	var detailBeer db.DetailBeer
	
	if err := c.ShouldBindJSON(&result); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// uploadImg(&result,c)
	nameBeer.Name = result.Name
	nameBeer.Type = result.Type
	nameBeer.Picture = result.Picture
	detailBeer.Type = result.Type
	detailBeer.Detail = result.Detail
	
	insertName := db.Db.Create(&nameBeer)
	if err := insertName.Error; err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"status": "cannot insert",
			"data":nameBeer,
		})
		return
	}
	
	

	insertDetail := db.Db.Create(&detailBeer)

	if err := insertDetail.Error; err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Status": "success",
		"dataDetail":detailBeer,
		"dataName":nameBeer,
	})
}


func uploadImg(data *db.Result,c *gin.Context)  {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		
	}
	now := time.Now().UnixNano()
	filename := strconv.FormatInt(now, 10) + ".jpg"
	
	if err := c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.Picture = filename
	c.JSON(http.StatusOK, gin.H{"status": "File uploaded successfully"})
}

