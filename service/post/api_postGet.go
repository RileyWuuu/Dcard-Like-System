package post

import (
	"dcard/storage/mongo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func postGet(c *gin.Context) {
	var p Post
	pst := &Post{}
	if err := json.NewDecoder(c.Request.Body).Decode(pst); err != nil {
		fmt.Println(err)
	}
	PostCollection := mongo.GetMongo().Collection("Post")
	objectid, err := primitive.ObjectIDFromHex(pst.Id)
	fmt.Println(pst)

	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = PostCollection.FindOne(ctx, bson.D{{"_id", objectid}}).Decode(&p)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonResp, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Error happened in Json marshal. Err: %s", err)
	}
	c.Writer.Write(jsonResp)
	return
}
