package post

import (
	"dcard/storage/mongo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func likeAdded(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	post := &Post{}
	PostCollection := mongo.GetMongo().Collection("Post")
	err := json.NewDecoder(r.Body).Decode(post)
	objectid, err := primitive.ObjectIDFromHex(post.Id)
	err = PostCollection.FindOne(ctx, bson.D{{"_id", objectid}}).Decode(&post)
	result, err := PostCollection.UpdateOne(
		ctx,
		bson.M{"_id": objectid},
		bson.D{
			{"$set", bson.D{{"likes", (post.Likes + 1)}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Added like : %v \n", result.ModifiedCount)
}
