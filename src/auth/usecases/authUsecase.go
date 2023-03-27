package usecases

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/AkapongAlone/komgrip-test/models"
	"github.com/AkapongAlone/komgrip-test/src/auth/domains"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	authRepo domains.AuthRepositories
}

func NewAuthUseCase(repo domains.AuthRepositories) domains.AuthUseCase {
	return &authUseCase{
		authRepo: repo,
	}
}

func (t *authUseCase) Register(username string, password string) error {
	fmt.Println(username, password)
	err := t.authRepo.IsExistUser(username)
	if err != nil {
		return err
	}
	encyptPass, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	password = string(encyptPass)
	fmt.Println(username, password)
	err = t.authRepo.CreateUser(username, password)
	if err != nil {
		return err
	}
	return nil
}

func (t *authUseCase) Login(username string, password string) (string, string, error) {
	var member models.User
	err := t.authRepo.IsExistUser(username)
	if err == nil {
		return "", "", errors.New("user not Exists.")
	}
	err = t.authRepo.FindUser(&member, username)
	err = bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(password))
	if err == nil {
		accessToken, err := t.CreateToken(int64(member.ID), time.Now().Add(15*time.Hour), "access")
		fmt.Println("access", accessToken)
		if err != nil {
			return "", "", err
		}
		refreshToken, err := t.CreateToken(int64(member.ID), time.Now().Add(24*time.Hour), "refresh")
		fmt.Println("refresh", refreshToken)
		if err != nil {
			return "", "", err
		}
		return accessToken, refreshToken, nil
	} else {
		return "", "", errors.New("Password not correct.")
	}

}

func (t *authUseCase) CreateToken(userID int64, expiration time.Time, typeToken string) (string, error) {
	sign := os.Getenv("SECRET")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = expiration.Unix()
	claims["type"] = typeToken
	fmt.Println("sign", sign)

	tokenString, err := token.SignedString([]byte(sign))
	fmt.Println("token", token)
	fmt.Println("token", tokenString)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *authUseCase)RefreshToken(tokenString string) (string, string, error){
    // validate the token
	sign := os.Getenv("SECRET")
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(sign), nil
    })
    if err != nil {
        return "","",err
    }

    // generate a new token with a new expiration time
    claims := token.Claims.(jwt.MapClaims)
	fmt.Println("tokwnnn",claims["type"].(string))
	if claims["type"].(string) != "refresh" {
		return "","",errors.New("Invalid token type")
	}
    ID := claims["userID"].(float64)
    userID := int64(ID)
	accessToken, err := t.CreateToken((userID), time.Now().Add(15*time.Hour), "access")
		fmt.Println("access", accessToken)
		if err != nil {
			return "", "", err
		}
	refreshToken, err := t.CreateToken((userID), time.Now().Add(24*time.Hour), "refresh")
	fmt.Println("refresh", refreshToken)
	if err != nil {
			return "", "", err
		}
	return accessToken, refreshToken, nil
}

