package models

import (
	_"mime/multipart"
	"time"
)


type BeerDB struct {
	ID        int       `form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string    `form:"name"  `
	Type      string    `form:"type" `
	Picture   string    `form:"picture" ` //เก็บเป็นpathของรูปภาพ
	Detail    string    `form:"detail" `
	CreatedAt time.Time `form:"created_at"`
	UpdatedAt time.Time `form:"updated_at"`
}
