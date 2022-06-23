package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func comment(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	db := MongoConn()
	PostCollection = db.Collection("Post")
	CommentCollection = db.Collection("Comment")
	cmt := &Comment{}
	err := json.NewDecoder(r.Body).Decode(cmt)
	cmt.CommentDate = time.Now()
	result, err := CommentCollection.InsertOne(ctx, cmt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(fmt.Sprintf("%v", result.InsertedID)))
	return
}
