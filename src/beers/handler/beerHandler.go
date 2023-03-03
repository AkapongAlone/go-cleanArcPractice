package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AkapongAlone/komgrip-test/models"
	_ "github.com/AkapongAlone/komgrip-test/requests"
	"github.com/AkapongAlone/komgrip-test/src/beers/domains"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type BeerHandler struct {
	beerUseCase domains.BeerUseCase
}

func NewBeerHandler(usecase domains.BeerUseCase) *BeerHandler {
	return &BeerHandler{beerUseCase: usecase}
}

func (t *BeerHandler) GetBeer(c *gin.Context) {
	nameParam := c.Query("name")
	limitParam := c.Query("limit")
	offsetParam := c.Query("offset")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, err := t.beerUseCase.GetBeer(nameParam, limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (t *BeerHandler) PostBeer(c *gin.Context) {
	var beerInputData models.Beer
	if err := c.ShouldBind(&beerInputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "hello",
		})
		return
	}

	var beer models.BeerDB
	copier.Copy(&beer, beerInputData)
	// file, header, err := c.Request.FormFile("file")
	file, err := c.FormFile("picture")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	beer.Picture, err = t.beerUseCase.Upload(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	fmt.Println(beer)

	err = t.beerUseCase.PostBeer(beer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (t *BeerHandler) UpdateBeer(c *gin.Context) {
	var result models.Beer
	if err := c.ShouldBind(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var beer models.BeerDB
	copier.Copy(&beer, result)
	// file, header, err := c.Request.FormFile("file")
	file, err := c.FormFile("picture")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	beer.Picture, err = t.beerUseCase.Upload(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	idParam := c.Param("id")
	fmt.Println(idParam,"idParam")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	fmt.Println(id)
	err = t.beerUseCase.UpdateBeer(id, beer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (t *BeerHandler) DeleteBeer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	err = t.beerUseCase.DeleteBeer(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "Deleted",
	})
}
