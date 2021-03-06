package post

import (
	"dcard/storage/mongo"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	redisClient "dcard/storage/redis"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func postCreate(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	PostCollection := mongo.GetMongo().Collection("Post")
	pst := &Post{}
	if err := json.NewDecoder(r.Body).Decode(pst); err != nil {
		fmt.Println(err)
	}
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
	if err != nil {
		fmt.Println(err)
	}

	_, errr := redisClient.GetRedis().ZAdd("Posts", redis.Z{timestamp, PJson}).Result()
	ErrorCheck(errr)

	w.Write([]byte(fmt.Sprintf("%v", result.InsertedID)))
	return
}
