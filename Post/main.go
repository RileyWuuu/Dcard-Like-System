package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("SERVER STARTED ON: HTTP://LOCALHOST:8091")
	http.HandleFunc("/", Index)
	http.HandleFunc("/member_insert", Insert)
	http.HandleFunc("/Update", Update)
	http.HandleFunc("/member_delete", Delete)
	http.HandleFunc("/GetPost", GetPost)
	http.HandleFunc("/GetPosts", GetPosts)
	http.HandleFunc("/CreatePost", CreatePost)
	http.HandleFunc("/PostComment", PostComment)
	http.HandleFunc("/GetComments", GetComments)
	http.HandleFunc("/AddLike", AddLike)
	// http.HandleFunc("/RedisConn", RedisConn)
	// http.HandleFunc("/SendRequest", SendRequest)
	// http.HandleFunc("/Matching", Matching)
	http.ListenAndServe(":8091", nil)
}
