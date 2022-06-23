package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func postGet(w http.ResponseWriter, r *http.Request) {
	db := mongoConn()
	var p Post
	pst := &Post{}
	err := json.NewDecoder(r.Body).Decode(pst)
	PostCollection = db.Collection("Post")
	objectid, err := primitive.ObjectIDFromHex(pst.Id)
	fmt.Println(pst)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = PostCollection.FindOne(ctx, bson.D{{"_id", objectid}}).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonResp, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Error happened in Json marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
