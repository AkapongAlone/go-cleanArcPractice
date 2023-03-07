package repositories

import (
	"errors"
	"fmt"

	"github.com/AkapongAlone/komgrip-test/models"
	_ "github.com/AkapongAlone/komgrip-test/requests"
	"github.com/AkapongAlone/komgrip-test/src/beers/handler"
	"github.com/AkapongAlone/komgrip-test/src/beers/usecases"
	"gorm.io/gorm"
)

type beerRepositories struct {
	conn *gorm.DB
}

func NewBeerHandler(conn *gorm.DB) *handler.BeerHandler	{
	usecase := usecases.NewBeerUseCase(&beerRepositories{conn: conn})
	handle := handler.NewBeerHandler(usecase)
	return handle
}

func (t *beerRepositories) CountAllData(name string) (int) {
	var count int64
	// t.conn.Model(&models.BeerDB{}).Where(&models.BeerDB{Name:name}).Count(&count)
	t.conn.Model(&models.BeerDB{}).Count(&count)
	fmt.Println(count)
	return int(count)
}

func (t *beerRepositories) FindData(name string,limit int,offset int) ([]models.BeerDB,error) {
	var result []models.BeerDB
	t.conn.Where(&models.BeerDB{Name:name}).Limit(limit).Offset(offset).Find(&result)
	if result == nil {
		return result,errors.New("data not found")
	}
	fmt.Println(len(result))
	return result,nil
}

func (t *beerRepositories) InsertData(name models.BeerDB)(error){
	err := t.conn.Create(&name).Error
	if err != nil { 
		return err
	}
	
	return nil
}

func (t *beerRepositories)EditData(id int,name models.BeerDB)(error) {
	fmt.Println(id)
	var beforeEdit models.BeerDB
	t.conn.Where("id = ?", id).First(&beforeEdit)
	name.ID = id
	err := t.conn.Model(&beforeEdit).Updates(name).Error
	if err != nil {
		return err
	}
	
	return nil
}
func(t *beerRepositories) DeleteData(id int)(error){
	
	err := t.conn.Delete(&models.BeerDB{},id).Error
	if err != nil { 
		return err
	}
	
	return nil

}