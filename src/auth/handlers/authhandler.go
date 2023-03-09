package handlers

import (
	 "fmt"
	"net/http"

	"github.com/AkapongAlone/komgrip-test/models"
	"github.com/AkapongAlone/komgrip-test/requests"
	"github.com/AkapongAlone/komgrip-test/responses"
	"github.com/AkapongAlone/komgrip-test/src/auth/domains"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type AuthHandler struct {
	authUseCase domains.AuthUseCase
}

func NewAuthHandler(usecase domains.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: usecase}
}

func (t *AuthHandler) Register(c *gin.Context) {
	var user models.User
	var userRequest requests.User

	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "hello",
		})
		return
	}

	copier.Copy(&user, userRequest)
	err := t.authUseCase.Register(user.UserName, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Status": "success",
	})
}

func (t *AuthHandler) Login(c *gin.Context) {
	var login requests.LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("get reqqq",login)
	accessToken, refreshToken, err := t.authUseCase.Login(login.Name, login.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("respond",accessToken,refreshToken)
	var loginResponse responses.LoginResponse
	loginResponse.AccessToken = accessToken
	loginResponse.RefreshToken = refreshToken
	fmt.Println("respond",loginResponse)
	c.JSON(http.StatusOK, gin.H{
		"token": loginResponse,
	})
}
