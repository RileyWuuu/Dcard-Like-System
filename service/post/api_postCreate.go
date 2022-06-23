package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func postCreate(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	db := mongoConn()
	client := redisConn()
	PostCollection = db.Collection("Post")
	pst := &Post{}
	err := json.NewDecoder(r.Body).Decode(pst)
	pst.PostDate = time.Now()
	pst.Id = ""
	now := time.Now()
	timestamp := float64(now.Unix())
	result, err := PostCollection.InsertOne(ctx, pst)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ContentTune := []rune(pst.Content)
	if len(pst.Content) > 30 {
		pst.Content = string(ContentTune[:30])
	}
	resultid := result.InsertedID
	ID := resultid.(primitive.ObjectID).Hex()

	post := map[string]string{
		"ID":      ID,
		"Title":   pst.Title,
		"Content": pst.Content,
		"Likes":   strconv.Itoa(pst.Likes),
	}
	PJson, err := json.Marshal(post)
	postString := PJson
	fmt.Println("postStringpostString", postString)

	_, errr := client.ZAdd("Posts", redis.Z{timestamp, PJson}).Result()
	ErrorCheck(errr)

	w.Write([]byte(fmt.Sprintf("%v", result.InsertedID)))
	return
}
