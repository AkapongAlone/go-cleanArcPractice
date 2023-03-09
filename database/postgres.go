package database

import (
	"fmt"
	"log"
	"os"
	_ "strconv"

	// "github.com/jinzhu/gorm"
	"github.com/AkapongAlone/komgrip-test/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

// BuildDBConfig use for building DB config
func BuildDBConfig() *DBConfig {
	// post, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE"),
	}
	return &dbConfig
}

// DbURL use for create DB connection URL
// func DbURL(dbConfig *DBConfig) string {
// 	return fmt.Sprintf(
// 		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
// 		dbConfig.User,
// 		dbConfig.Password,
// 		dbConfig.Host,
// 		dbConfig.Port,
// 		dbConfig.DBName,
// 	)
// }

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln("consider env var")
	} 
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE"),
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",dbConfig.Host,dbConfig.User,dbConfig.Password,dbConfig.DBName,dbConfig.Port)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	Db.AutoMigrate(&models.BeerDB{})
	Db.AutoMigrate(&models.User{})
	fmt.Println("Db connect")
}