package post

import (
	"context"
	"database/sql"
	"log"
)

func mysqlConn() (db *sql.DB) {
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

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
