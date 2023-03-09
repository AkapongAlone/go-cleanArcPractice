package repositories

import (
	"errors"
	"fmt"

	"github.com/AkapongAlone/komgrip-test/models"
	"github.com/AkapongAlone/komgrip-test/src/auth/handlers"
	"github.com/AkapongAlone/komgrip-test/src/auth/usecases"
	"gorm.io/gorm"
)

type authRepositories struct {
	conn *gorm.DB
}

func NewAuthHandler(conn *gorm.DB) *handlers.AuthHandler {
	usecase := usecases.NewAuthUseCase(&authRepositories{conn: conn})
	handle := handlers.NewAuthHandler(usecase)
	return handle
}

func (t *authRepositories) CreateUser(username string,password string)(error){
	var userBody models.User
	userBody.UserName = username
	userBody.Password = password
	err := t.IsExistUser(username)
	if err != nil {
		return err
	}
	err = t.conn.Create(&userBody).Error
	
	if err != nil {
		return err
	}
	return nil
}

func (t *authRepositories) IsExistUser(name string)(error){
	var data models.User
	count := int64(0)
	fmt.Println("is existttttt",name)
	if name != "" {
		err := t.conn.Model(&data).Where("user_name = ?", name).Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("Already have this name.")
		}
	}
	return nil
}

func (t *authRepositories) FindUser(model *models.User,name string) error {
	err := t.conn.Where("user_name = ?", name).First(&model).Error
	if err != nil {
		return err
	}
	return nil
}