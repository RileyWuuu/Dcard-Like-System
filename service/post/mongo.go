package post

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	CommentCollection *mongo.Collection
	PostCollection    *mongo.Collection
	ctx               = context.TODO()
)

func mongoConn() (db *mongo.Database) {
	host := "127.0.0.1"
	port := "27017"
	connectionURI := "mongodb://" + host + ":" + port + "/"
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	mdb := client.Database("testdb")
	PostCollection = mdb.Collection("Post")
	CommentCollection = mdb.Collection("Comment")
	return mdb
}
