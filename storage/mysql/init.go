package mysql

import (
	"database/sql"
	"log"
	"sync"
)

var db *sql.DB
var once sync.Once
var initialized bool

func GetMySQL() *sql.DB {
	if !initialized {
		Initialize()
	}

	return db
}

func Initialize() {
	once.Do(func() {
		dbDriver := "mysql"
		dbUser := "root"
		dbPass := "0000"
		dbName := "testdb"
		c, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
		if err != nil {
			log.Fatal(err)
		}

		db = c
		initialized = true
	})
}
