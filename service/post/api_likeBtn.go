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
	post := &Post{}
	PostCollection := mongo.GetMongo().Collection("Post")
	if err := json.NewDecoder(r.Body).Decode(post); err != nil {
		fmt.Println(err)
	}
	memID := post.MemberID
	objectid, err := primitive.ObjectIDFromHex(post.Id)
	if err != nil {
		fmt.Println(err)
	}
	err = PostCollection.FindOne(ctx, bson.D{{"_id", objectid}}).Decode(&post)
	checkIfLiked := false
	for _, x := range post.Liked {
		if x == memID {
			checkIfLiked = true
		}
	}
	if checkIfLiked == true {
		result, err := PostCollection.UpdateOne(
			ctx,
			bson.M{"_id": objectid},
			bson.D{
				{"$set", bson.D{{"likes", (post.Likes - 1)}}},
				{"$pull", bson.D{{"liked", memID}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Printf("Added like : %v \n", result.ModifiedCount)
		return
	} else {
		result, err := PostCollection.UpdateOne(
			ctx,
			bson.M{"_id": objectid},
			bson.D{
				{"$set", bson.D{{"likes", (post.Likes + 1)}}},
				{"$push", bson.D{{"liked", memID}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Printf("Added like : %v \n", result.ModifiedCount)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return

}
