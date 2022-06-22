package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
)

func MysqlConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "0000"
	dbName := "testdb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

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
