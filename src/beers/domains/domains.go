package domains

import (
	"mime/multipart"

	"github.com/AkapongAlone/komgrip-test/models"
	_ "github.com/AkapongAlone/komgrip-test/requests"
)

type BeerUseCase interface {
	GetBeer(name string,limit int,offset int) ([]models.BeerDB,error)
	PostBeer(request models.BeerDB) (error)
	UpdateBeer(id int,request models.BeerDB) (error)
	DeleteBeer(id int) (error)
	Upload(header *multipart.FileHeader)(string,error)
}

type BeerRepositories interface {
	FindData(name string,limit int,offset int) ([]models.BeerDB,error)
	InsertData(data models.BeerDB)(error)
	EditData(id int,data models.BeerDB)(error)
	DeleteData(id int)(error)
	
}