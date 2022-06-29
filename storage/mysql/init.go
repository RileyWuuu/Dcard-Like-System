package mysql

import (
	"database/sql"
	"dcard/config"
	"fmt"
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
		addr := fmt.Sprintf("%s:%s@/%s", config.Conf.MySql.UserName, config.Conf.MySql.Password, config.Conf.MySql.DbName)
		c, err := sql.Open(config.Conf.MySql.DriverName, addr)
		if err != nil {
			log.Fatal(err)
		}

		db = c
		initialized = true
	})
}
