package orm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error
type NameBeer struct {
	gorm.Model
	Name	string `json:"name"  `
	Type 	string `json:"type" `
	Picture	string `json:"picture" `  //เก็บเป็นpathของรูปภาพ
}

type DetailBeer struct {
	gorm.Model
	Type 	string `json:"type"  `
	Detail	string `json:"detail" `
}

type Result struct {
	Name  string `json:"name"`
	Type string `json:"type"`
	Picture	string `json:"picture"`
	Detail string `json:"detail"`
  }


func InitDB() {
	dsn := "host=localhost  user=postgres password=1234  dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	Db.AutoMigrate(&NameBeer{})
	Db.AutoMigrate(&DetailBeer{})

	fmt.Println("Db connect")
}