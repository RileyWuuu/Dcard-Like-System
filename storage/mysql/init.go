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
		addr := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Conf.MySql.UserName, config.Conf.MySql.Password, config.Conf.MySql.Addr, config.Conf.MySql.DbName)
		fmt.Println("AAAAAAA", addr)
		c, err := sql.Open(config.Conf.MySql.DriverName, addr)
		if err != nil {
			log.Fatal(err)
		}

		db = c
		initialized = true
	})
}
