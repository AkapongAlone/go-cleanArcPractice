package domains

import "github.com/AkapongAlone/komgrip-test/models"

type AuthUseCase interface {
	Register(username string, password string) error
	Login(username string, password string) (string, string, error)
	RefreshToken(tokenString string) (string, string, error)
}

type AuthRepositories interface {
	CreateUser(username string, password string) error
	IsExistUser(username string) error
	FindUser(model *models.User,name string) error
}