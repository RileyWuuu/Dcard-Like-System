package matching

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func EnableMatchingServer() {
	log.Println("SERVER STARTED ON: HTTP://LOCALHOST:8092")
	// http.HandleFunc("/RedisConn", RedisConn)
	// http.HandleFunc("/SendRequest", SendRequest)
	http.HandleFunc("/Matching", matching)
	http.ListenAndServe(":8092", nil)
}
