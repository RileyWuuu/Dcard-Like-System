package post

import (
	"dcard/storage/mongo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func commentsGet(c *gin.Context) {
	var condition bson.D
	cmt := &Comment{}
	var comment Comment
	var comments []Comment
	if err := json.NewDecoder(c.Request.Body).Decode(cmt); err != nil {
		fmt.Println(err)
	}
	condition = append(condition, bson.E{Key: "postid", Value: cmt.PostID})
	fmt.Println(cmt)
	CommentCollection := mongo.GetMongo().Collection("Comment")
	cursor, err := CommentCollection.Find(ctx, condition)
	if err != nil {
		defer cursor.Close(ctx)
		fmt.Println("ERROR")
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(&comment)
		fmt.Println(err)
		if err != nil {
			fmt.Println("ERROR")
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		comments = append(comments, comment)
	}
	jsonResp, err := json.Marshal(comments)
	if err != nil {
		log.Fatalf("Error happened in Json marshal. Err: %s", err)
	}
	c.Writer.Write(jsonResp)
	return
}
