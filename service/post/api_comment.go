package post

import (
	"dcard/storage/mongo"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func comment(c *gin.Context) {
	// Authentication(w, r)
	CommentCollection := mongo.GetMongo().Collection("Comment")
	cmt := &Comment{}
	if err := json.NewDecoder(c.Request.Body).Decode(cmt); err != nil {
		fmt.Println(err)
	}

	cmt.CommentDate = time.Now()
	result, err := CommentCollection.InsertOne(ctx, cmt)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	c.Writer.Write([]byte(fmt.Sprintf("%v", result.InsertedID)))
	return
}
