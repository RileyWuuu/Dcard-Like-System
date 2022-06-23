package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func postsGet(w http.ResponseWriter, r *http.Request) {
	var post PostSummary
	var posts []PostSummary
	var posts2 []PostSummary
	db := mongoConn()
	rdb := redisConn()
	page := &Posts{}
	err := json.NewDecoder(r.Body).Decode(page)
	ErrorCheck(err)
	fmt.Println(page)
	total := page.Page * page.PerPage
	result := rdb.ZRevRange("Posts", 0, int64(total))
	ErrorCheck(err)
	aaa := result.Val()
	for _, count := range aaa {
		arr := strings.Split(count, ",")

		post.Content = arr[0]
		post.Id = arr[1]
		post.Likes, _ = strconv.Atoi(arr[2])
		post.Title = arr[3]
		posts = append(posts, post)
	}

	if len(posts) != total {
		fmt.Println(len(posts), total)
		PostCollection = db.Collection("Post")
		findOptions := options.Find()
		findOptions.SetSort(bson.D{{"postdate", -1}})

		cursor, err := PostCollection.Find(ctx, bson.D{}, findOptions)
		if err != nil {
			defer cursor.Close(ctx)
			fmt.Println("ERROR")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for cursor.Next(ctx) {
			err := cursor.Decode(&post)
			if err != nil {
				fmt.Println("ERROR")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ContentTune := []rune(post.Content)
			if len(post.Content) > 50 {
				post.Content = string(ContentTune[:50])
			}
			posts2 = append(posts2, post)
		}
	}
	posts2 = posts2[len(posts)-1:]
	posts2 = posts2[:total-len(posts)]
	for _, data := range posts2 {
		posts = append(posts, data)
	}
	jsonResp, err := json.Marshal(posts2)
	if err != nil {
		log.Fatalf("Error happened in Json marshal. Err: %s", err)
	}
	fmt.Println("posts2:", len(posts))

	w.Write(jsonResp)
	return
}
