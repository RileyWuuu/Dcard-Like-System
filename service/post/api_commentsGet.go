package post

import (
	"dcard/storage/mongo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func commentsGet(w http.ResponseWriter, r *http.Request) {
	var condition bson.D
	cmt := &Comment{}
	var comment Comment
	var comments []Comment
	err := json.NewDecoder(r.Body).Decode(cmt)
	condition = append(condition, bson.E{Key: "postid", Value: cmt.PostID})
	fmt.Println(cmt)
	CommentCollection := mongo.GetMongo().Collection("Comment")
	cursor, err := CommentCollection.Find(ctx, condition)
	if err != nil {
		defer cursor.Close(ctx)
		fmt.Println("ERROR")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(&comment)
		fmt.Println(err)
		if err != nil {
			fmt.Println("ERROR")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		comments = append(comments, comment)
	}
	jsonResp, err := json.Marshal(comments)
	if err != nil {
		log.Fatalf("Error happened in Json marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
