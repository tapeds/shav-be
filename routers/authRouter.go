package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapeds/shav-be/controllers"
)

func Router(route *gin.Engine){
	routes:=route.Group("/api/user")
	{
		routes.POST("/login", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{"data": "login"})
		})
		routes.POST("", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{"data": "register"})
		})
		routes.GET("", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{"data": "get user"})
		})
	}
}