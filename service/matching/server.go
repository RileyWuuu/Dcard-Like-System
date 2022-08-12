package matching

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func EnableMatchingServer() {
	server := gin.Default()
	server.GET("/Matching", matching)
	server.Run("localhost:8092")
}
