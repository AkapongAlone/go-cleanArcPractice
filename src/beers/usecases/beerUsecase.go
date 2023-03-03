package usecases

import (
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/AkapongAlone/komgrip-test/models"
	_ "github.com/AkapongAlone/komgrip-test/requests"
	"github.com/AkapongAlone/komgrip-test/src/beers/domains"
)

type beerUseCase struct {
	beerRepo domains.BeerRepositories
}

func NewBeerUseCase(repo domains.BeerRepositories) domains.BeerUseCase {
	return &beerUseCase{
		beerRepo: repo,
	}
}

func (t *beerUseCase)GetBeer(name string,limit int,offset int) ([]models.BeerDB,error){
		result,err := t.beerRepo.FindData(name,limit,offset)
		if err != nil {
			return result,err
		}
		return result,nil
}

func (t *beerUseCase)PostBeer(result models.BeerDB) (error){
		
		err := t.beerRepo.InsertData(result)
		if err != nil {
			return err
		}
		return nil
	
}

func (t *beerUseCase)UpdateBeer(id int,result models.BeerDB) (error){
	// var nameBeer models.BeerDB
	
	err := t.beerRepo.EditData(id,result)
	if err != nil {
		return err
	}
	return nil
}

func (t *beerUseCase)DeleteBeer(id int) (error){
	err := t.beerRepo.DeleteData(id)
	if err != nil {
		return err
	}
	return nil
}

func (t *beerUseCase)Upload(header *multipart.FileHeader)(string,error){
	
	fileName:= header.Filename
	src, err := header.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	out, err := os.Create("./images/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	if _, err := io.Copy(out, src); err != nil {
		return "", err
	}
	filePath := "http://localhost:8080/file/" + fileName
	return filePath,nil
}