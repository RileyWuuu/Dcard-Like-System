package apigateway

import (
	"github.com/gin-gonic/gin"
)

func EnableApiGateway() {
	server := gin.Default()
	server.POST("/Login", login)
	server.POST("/Refresh", refresh)
	server.Run("localhost:8090")
}
