package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("SERVER STARTED ON: HTTP://LOCALHOST:8092")
	// http.HandleFunc("/RedisConn", RedisConn)
	// http.HandleFunc("/SendRequest", SendRequest)
	http.HandleFunc("/Matching", Matching)
	http.ListenAndServe(":8092", nil)
}
