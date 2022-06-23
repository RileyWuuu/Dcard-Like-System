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
	total = total - 1
	result := rdb.ZRevRange("Posts", 0, int64(total))
	ErrorCheck(err)
	aaa := result.Val()
	i := 0
	for _, count := range aaa {
		i = i + 1
		arr := strings.Split(count, ",")

		post.Content = arr[0]
		post.Id = arr[1]
		post.Likes, _ = strconv.Atoi(arr[2])
		post.Title = arr[3]
		posts = append(posts, post)
	}

	if i != total {
		fmt.Println(i, total)
		// k := total - i
		PostCollection = db.Collection("Post")
		findOptions := options.Find()
		findOptions.SetSort(bson.D{{"postdate", -1}})

		// opts := options.Find().SetSort(bson.D{{"postdate", -1}})
		cursor, err := PostCollection.Find(ctx, bson.D{}, findOptions)
		//cursor := cursor.SetSort(bson.D{{"_id",-1}})
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
	posts2 = append(posts2[i+1])
	jsonResp, err := json.Marshal(posts2)
	if err != nil {
		log.Fatalf("Error happened in Json marshal. Err: %s", err)
	}
	// difference := total - i
	fmt.Println("posts2:", posts2)

	w.Write(jsonResp)
	return
}
