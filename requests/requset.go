package requests

import (
	"mime/multipart"
	"time"
)

type Beer struct {
	// ID        int       `form:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string                `form:"name"  `
	Type      string                `form:"type" `
	Picture   *multipart.FileHeader `form:"picture" ` //เก็บเป็นpathของรูปภาพ
	Detail    string                `form:"detail" `
	CreatedAt time.Time             `form:"created_at"`
	UpdatedAt time.Time             `form:"updated_at"`
}

type User struct {
	UserName	string 		`form:"name"  `
	Password	string		`form:"password"  `
}

type LoginRequest struct {
	Name    string `json:"name"`
	Password string `json:"password" `
}