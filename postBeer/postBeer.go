package postbeer

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	db "github.com/AkapongAlone/komgrip-test/orm"
	"github.com/gin-gonic/gin"
)

type result struct {
	Name  string
	Type string
	Picture	string
	Detail string
  }

func PostBeer(c *gin.Context) {
	var result result
	var nameBeer db.NameBeer
	
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
	fmt.Println(nameBeer)
	insertName := db.Db.Create(&nameBeer)
	if err := insertName.Error; err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	var detailBeer db.DetailBeer
	detailBeer.Type = result.Type
	detailBeer.Type = result.Detail

	insertDetail := db.Db.Create(&detailBeer)

	if err := insertDetail.Error; err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Status": "success",
		
	})
}


func uploadImg(data *result,c *gin.Context)  {
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

