package router

import (
	"github.com/go-whisper/go-whisper/app/controller/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-whisper/go-whisper/app/instance"
)

func init() {
	service := instance.WebService()
	// default:
	{
		frontend := service.Group("")
		{
			ctr := web.Index{}
			frontend.GET("", ctr.Index)
		}
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
