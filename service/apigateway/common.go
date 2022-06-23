package apigateway

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
)

var jwtKey = []byte("Secret")

var ctxb = context.Background()
var (
	CommentCollection *mongo.Collection
	PostCollection    *mongo.Collection
	ctx               = context.TODO()
)

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
