package orm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type NameBeer struct {
	gorm.Model
	Name	string `json:"name" binding:"required" `
	Type 	string `json:"type" binding:"required"`
	Picture	string `json:"picture" binding:"required"`  //เก็บเป็นpathของรูปภาพ
}

type DetailBeer struct {
	gorm.Model
	Type 	string `json:"type" binding:"required" gorm:"primaryKey"`
	Detail	string `json:"detail" binding:"required"`
}

func InitDB() {
	dsn := "host=localhost  user=postgres password=1234  dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Db.AutoMigrate(&NameBeer{})
	Db.AutoMigrate(&DetailBeer{})
	fmt.Println("Db connect")
}