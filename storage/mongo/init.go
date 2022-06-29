package mongo

import (
	"context"
	"dcard/config"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var once sync.Once
var initialized bool

func GetMongo() *mongo.Database {
	if !initialized {
		Initialize()
	}

	return db
}

func GetCollection(colName string) *mongo.Collection {
	if !initialized {
		Initialize()
	}

	return db.Collection(colName)
}

func Initialize() {
	once.Do(func() {
		connectionURI := fmt.Sprintf("mongodb://%s/", config.Conf.Mongo.Addr)
		clientOptions := options.Client().ApplyURI(connectionURI)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}

		db = client.Database("testdb")
		initialized = true
	})
}
