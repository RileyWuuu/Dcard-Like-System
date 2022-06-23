package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("SERVER STARTED ON: HTTP://LOCALHOST:8091")
	http.HandleFunc("/", membersGet)
	http.HandleFunc("/member_insert", Insert)
	http.HandleFunc("/Update", Update)
	http.HandleFunc("/member_delete", Delete)
	http.HandleFunc("/postGet", postGet)
	http.HandleFunc("/postsGet", postsGet)
	http.HandleFunc("/postCreate", postCreate)
	http.HandleFunc("/comment", comment)
	http.HandleFunc("/commentsGet", commentsGet)
	http.HandleFunc("/likeAdded", likeAdded)
	http.ListenAndServe(":8091", nil)
}
