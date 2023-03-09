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

func NewBeerHandler(conn *gorm.DB) *handler.BeerHandler {
	usecase := usecases.NewBeerUseCase(&beerRepositories{conn: conn})
	handle := handler.NewBeerHandler(usecase)
	return handle
}

func (t *beerRepositories) CountAllData(name string) int {
	var count int64
	// t.conn.Model(&models.BeerDB{}).Where(&models.BeerDB{Name:name}).Count(&count)
	t.conn.Model(&models.BeerDB{}).Count(&count)
	fmt.Println(count)
	return int(count)
}

func (t *beerRepositories) FindData(name string, limit int, offset int) ([]models.BeerDB, error) {
	var result []models.BeerDB
	query := t.conn.Where(&models.BeerDB{Name: name}).Offset(offset).Find(&result)
	if result == nil {
		return result, errors.New("data not found")
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	query.Find(&result)
	return result, nil
}

func (t *beerRepositories) InsertData(name models.BeerDB) error {
	err := t.IsExistByName(name.Name)
	if err != nil {
		return err
	}
	err = t.conn.Create(&name).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *beerRepositories) EditData(newData *models.BeerDB) error {
	// var beforeEdit models.BeerDB

	err := t.conn.Save(newData).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *beerRepositories) FindByID(id int) (models.BeerDB, error) {
	var beforeEdit models.BeerDB
	err := t.conn.Where("id = ?", id).First(&beforeEdit).Error
	if err != nil {
		return beforeEdit, err
	}
	return beforeEdit, nil
}

func (t *beerRepositories) DeleteData(id int) error {

	_,err := t.FindByID(id)
	if err != nil {
		return err
	}
	err = t.conn.Delete(&models.BeerDB{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *beerRepositories) FindPicture(id int) (string, error) {

	var data models.BeerDB
	err := t.conn.Where("id = ?", id).First(&data).Error
	if err != nil {
		return "", err
	}
	return data.Picture, nil
}


func (t *beerRepositories) IsExistByName(name string) error {
	var data models.BeerDB
	count := int64(0)
	if name != "" {
		err := t.conn.Model(&data).Where("name = ?", name).Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("Already have this name.")
		}
	}
	return nil
}
