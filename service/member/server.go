package member

import (
	"github.com/gin-gonic/gin"
)

func EnableMemberServer() {
	server := gin.Default()
	server.GET("/member_get", singleMemberGet)
	server.POST("/member_delete", delete)
	server.POST("/member_insert", insert)
	server.POST("/member_update", update)
	server.GET("/get", membersGet)
	server.Run("localhost:8093")
}
