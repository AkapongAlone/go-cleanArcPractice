package domains

import (
	"mime/multipart"

	"github.com/AkapongAlone/komgrip-test/models"
	_ "github.com/AkapongAlone/komgrip-test/requests"
	"github.com/AkapongAlone/komgrip-test/responses"

)

type BeerUseCase interface {
	GetBeer(name string,limit int,offset int) (responses.PaginationBody,error)
	GetBeerById(id int) (models.BeerDB,error)
	PostBeer(request models.BeerDB) (error)
	UpdateBeer(id int,request *models.BeerDB) (error)
	DeleteBeer(id int) (error)
	Upload(header *multipart.FileHeader)(string,error)
	Remove(id int)(error)
	GetImg(id int) (string,error)
}

type BeerRepositories interface {
	FindData(name string,limit int,offset int) ([]models.BeerDB,error)
	InsertData(data models.BeerDB)(error)
	EditData(data *models.BeerDB)(error)
	DeleteData(id int)(error)
	CountAllData(name string)(int)
	FindPicture(id int)(string,error)
	FindByID(id int) (models.BeerDB,error)
	IsExistByName(name string) (error)
}