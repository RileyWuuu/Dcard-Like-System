package post

import (
	"log"
	"net/http"
)

func EnablePostServer() {
	log.Println("SERVER STARTED ON: HTTP://LOCALHOST:8091")
	// http.HandleFunc("/", index)
	http.HandleFunc("/member_insert", insert)
	http.HandleFunc("/Update", update)
	http.HandleFunc("/member_delete", delete)
	http.HandleFunc("/GetPost", postGet)
	http.HandleFunc("/GetPosts", postsGet)
	http.HandleFunc("/CreatePost", postCreate)
	http.HandleFunc("/PostComment", comment)
	http.HandleFunc("/GetComments", commentsGet)
	http.HandleFunc("/AddLike", likeAdded)
	http.ListenAndServe(":8091", nil)
}
