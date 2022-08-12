package post

import (
	"dcard/storage/mongo"
	"dcard/storage/redis"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func postsGet(c *gin.Context) {
	var post PostSummary
	var posts []PostSummary
	var posts2 []PostSummary
	page := &Posts{}
	err := json.NewDecoder(c.Request.Body).Decode(page)
	ErrorCheck(err)
	total := page.Page * page.PerPage

	result := redis.GetRedis().ZRevRange("Posts", 0, int64(total))
	ErrorCheck(err)
	aaa := result.Val()
	PostCollection := mongo.GetMongo().Collection("Post")
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"postdate", -1}})
	//If redis has no data, get all from mongo
	if len(aaa) == 0 {
		cursor, err := PostCollection.Find(ctx, bson.D{}, findOptions)
		if err != nil {
			defer cursor.Close(ctx)
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		for cursor.Next(ctx) {
			err := cursor.Decode(&post)
			if err != nil {
				c.Writer.WriteHeader(http.StatusBadRequest)
				return
			}
			ContentTune := []rune(post.Content)
			if len(post.Content) > 50 {
				post.Content = string(ContentTune[:50])
			}
			posts2 = append(posts2, post)
		}

		posts2 = posts2[:total]
		for _, data := range posts2 {
			posts = append(posts, data)
		}
		jsonResp, err := json.Marshal(posts)
		if err != nil {
			log.Fatalf("Error happened in Json marshal. Err: %s", err)
		}
		c.Writer.Write(jsonResp)

	} else { //
		fmt.Println(len(aaa))
		for _, count := range aaa {
			arr := strings.Split(count, ",")
			post.Content = arr[0]
			post.Id = arr[1]
			post.Likes, _ = strconv.Atoi(arr[2])
			post.Title = arr[3]
			posts = append(posts, post)
		}
		fmt.Println("len", len(posts), "total:", total)
		if len(posts) < total {
			cursor, err := PostCollection.Find(ctx, bson.D{}, findOptions)
			if err != nil {
				defer cursor.Close(ctx)
				c.Writer.WriteHeader(http.StatusBadRequest)
				return
			}
			for cursor.Next(ctx) {

				err := cursor.Decode(&post)
				if err != nil {
					c.Writer.WriteHeader(http.StatusBadRequest)
					return
				}
				ContentTune := []rune(post.Content)
				if len(post.Content) > 50 {
					post.Content = string(ContentTune[:50])
				}
				posts2 = append(posts2, post)
			}
		}
		posts2 = posts2[:total-len(posts)]
		for _, data := range posts2 {
			posts = append(posts, data)
		}
		jsonResp, err := json.Marshal(posts)
		if err != nil {
			log.Fatalf("Error happened in Json marshal. Err: %s", err)
		}
		c.Writer.Write(jsonResp)
	}
	fmt.Println(posts)

	return
}
