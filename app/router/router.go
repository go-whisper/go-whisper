package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-whisper/go-whisper/app/instance"
)

func init() {
	service := instance.WebService()
	// web:
	{
		frontend := service.Group("")
		frontend.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "首页")
		})
	}
	{
		admin := service.Group("/admin")
		admin.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "管理-首页")
		})
	}

	// api:
	api := service.Group("/api")
	api.GET("version", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"version": "0.0.1"})
	})
}
