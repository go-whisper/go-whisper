package router

import (
	"github.com/go-whisper/go-whisper/app/controller/middleware"
	"github.com/go-whisper/go-whisper/app/controller/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-whisper/go-whisper/app/instance"
)

func init() {
	service := instance.WebService()
	service.Use(middleware.Tracker(), middleware.InitUser())
	// default:
	{
		frontend := service.Group("")
		{
			ctr := web.Index{}
			frontend.GET("", ctr.Index)
		}
		{
			ctr := web.User{}
			frontend.GET("login", ctr.LoginForm)
			frontend.POST("login", ctr.Login)
			frontend.POST("reset-password", ctr.ResetPassword)
		}
		{
			ctr := web.Post{}
			group := frontend.Group("posts")
			group.GET(":id/delete", ctr.Remove)
		}
	}

	// api:
	api := service.Group("/api")
	api.GET("version", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"version": "0.0.1"})
	})
}
