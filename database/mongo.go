package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var collection *mongo.Collection

const (
	DBNAME  = "loggerAPI"
	COLNAME = "logInfo"
)

func InitMONGO() {
	mongoURI := os.Getenv("MONGO_URI")
	clientOption := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatalln(err.Error())
	}

	collection = client.Database(DBNAME).Collection(COLNAME)
}

func GetMONGO() *mongo.Collection {
	return collection
}