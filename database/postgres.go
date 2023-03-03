package database

import (
	"fmt"
	"os"
	"strconv"
	// "github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/AkapongAlone/komgrip-test/models"
)

var Db *gorm.DB
var err error
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig use for building DB config
func BuildDBConfig() *DBConfig {
	post, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     post,
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE"),
	}
	return &dbConfig
}

// DbURL use for create DB connection URL
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func InitDB() {
	dsn := "host=localhost  user=postgres password=1234  dbname=test_aek port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	Db.AutoMigrate(&models.BeerDB{})

	fmt.Println("Db connect")
}