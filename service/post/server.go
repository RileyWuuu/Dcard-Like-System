package post

import (
	"github.com/gin-gonic/gin"
)

func EnablePostServer() {
	server := gin.Default()
	server.GET("/GetPost", postGet)
	server.GET("/GetPosts", postsGet)
	server.GET("/GetComments", commentsGet)
	server.POST("/AddLike", likeAdded)
	server.POST("/CreatePost", postCreate)
	server.POST("/PostComment", comment)
	server.Run("localhost:8093")
}
