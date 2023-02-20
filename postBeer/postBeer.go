package postbeer

import (
	"net/http"
	"strconv"
	"time"

	db "github.com/AkapongAlone/komgrip-test/orm"
	"github.com/gin-gonic/gin"
)

func PostBeer(c *gin.Context) {
	var nameBeer db.NameBeer
	
	if err := c.ShouldBindJSON(&nameBeer); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	uploadImg(&nameBeer,c)
	insertName := db.Db.Create(&nameBeer)
	if err := insertName.Error; err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"Status": "success",
		
	})
}

func PostDetail(c *gin.Context) {
	var detailBeer db.DetailBeer
	if err := c.ShouldBindJSON(&detailBeer); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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
		
	})
}

func uploadImg(data *db.NameBeer,c *gin.Context)  {
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

