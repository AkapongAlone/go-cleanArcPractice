package usecases

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/AkapongAlone/komgrip-test/helper"
	"github.com/AkapongAlone/komgrip-test/models"
	_ "github.com/AkapongAlone/komgrip-test/requests"
	"github.com/AkapongAlone/komgrip-test/responses"
	"github.com/AkapongAlone/komgrip-test/src/beers/domains"
	"github.com/jinzhu/copier"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

type beerUseCase struct {
	beerRepo domains.BeerRepositories
}

func NewBeerUseCase(repo domains.BeerRepositories) domains.BeerUseCase {
	return &beerUseCase{
		beerRepo: repo,
	}
}

func (t *beerUseCase) GetBeerById(id int) (models.BeerDB,error) {
	result,err := t.beerRepo.FindByID(id)
	if err != nil {
		return result,err
	}
	return result,nil
}

func (t *beerUseCase) GetImg(id int) (string,error) {
	result,err := t.beerRepo.FindByID(id)
	if err != nil {
		return "",err
	}
	return result.Picture,nil
}

func (t *beerUseCase) GetBeer(name string, limit int, page int) (responses.PaginationBody, error) {
	
	totalItem := t.beerRepo.CountAllData(name)
	var totalPage int
	if limit != 0 {
		totalPage = totalItem / limit
		fmt.Println("totalpage1",totalPage)
		switch {
		case totalPage == 0:
			totalPage = 1
		case (totalItem%limit != 0):
			totalPage++
		}
	} else {
		totalPage = 1
		limit = totalItem
	}
	fmt.Println("totalpage2",totalPage)
	switch {
	case page <= 0 :page =1
	case page > totalPage:page = totalPage
	}

	var items []responses.ItemBody
	offset := (page - 1) * limit
	result, err := t.beerRepo.FindData(name, limit, offset)
	if err != nil {
		return responses.PaginationBody{}, err
	}
	for _, item := range result {
		itemBody := responses.ItemBody{
			Created_at: item.CreatedAt.String(),
			ID:         item.ID,
			Name:       name,
			Picture:    helper.GetPathImg(item.Picture),
			Detail:     item.Detail,
			Type:       item.Type,
			Updated_at: item.UpdatedAt.String(),
		}
		items = append(items, itemBody)
	}

	var nextPage, prevPage int
	if page == totalPage {
		nextPage = page
	} else {
		nextPage = page + 1
	}
	if page > 1 {
		prevPage = page - 1
	} else {
		prevPage = page
	}
	respond := responses.PaginationBody{
		CurrentPage:  page,
		Items:        items,
		NextPage:     nextPage,
		PreviousPage: prevPage,
		SizePerPage:  limit,
		TotalItems:   totalItem,
		TotalPages:   totalPage,
	}
	return respond, nil
}

func (t *beerUseCase) PostBeer(result models.BeerDB) error {

	err := t.beerRepo.InsertData(result)
	if err != nil {
		return err
	}
	return nil

}

func (t *beerUseCase) UpdateBeer(id int, newData *models.BeerDB) (error) {
	// var oldData models.BeerDB
	oldData,err := t.beerRepo.FindByID(id)
	if err != nil {
		return err
	}
	if oldData.Name != newData.Name {
		err = t.beerRepo.IsExistByName(newData.Name)
		if err != nil {
			return err
		}
	}
	

	
	oldData.Name = newData.Name
	oldData.Type = newData.Type
	oldData.Picture = newData.Picture
	oldData.Detail = newData.Detail
	err = t.beerRepo.EditData(&oldData)

	copier.Copy(&newData, oldData)
	// err = t.beerRepo.FindByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (t *beerUseCase) DeleteBeer(id int) error {
	err := t.beerRepo.DeleteData(id)
	if err != nil {
		return err
	}
	return nil
}

func (t *beerUseCase) Upload(header *multipart.FileHeader) (string, error) {

	fileName := header.Filename
	src, err := header.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	
	timestamp := time.Now().UnixNano()
    // create the new file name with the timestamp
    newFileName := strconv.FormatInt(timestamp, 10) + fileName
	out, err := os.Create("./images/" + newFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	if _, err := io.Copy(out, src); err != nil {
		return "", err
	}
	// filePath := "http://localhost:8080/file/" + fileName
	return newFileName, nil
}

func (t *beerUseCase) Remove(id int) error {

	fileName, err := t.beerRepo.FindPicture(id)
	fmt.Println(fileName)
	if err != nil {
		return err
	}
	err = os.Remove("./images/" + fileName)
	if err != nil {
		return err
	}
	return nil
}
