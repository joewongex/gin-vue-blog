package routes

import (
	"gvblog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	g := r.Group("/api/v1")
	{
		g.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "hello",
			})
		})
	}

	r.Run(utils.HttpPort)
}
