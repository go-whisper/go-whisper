package router

import (
	"net/http"

	"github.com/go-whisper/go-whisper/app/controller/middleware"
	"github.com/go-whisper/go-whisper/app/controller/web"

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
			ctr := web.Post{}
			frontend.GET("", ctr.Index)
			frontend.GET("posts/:id", ctr.Detail)
		}
		{
			ctr := web.User{}
			frontend.GET("users/login", ctr.LoginForm)
			frontend.POST("users/login-do", ctr.Login)
			frontend.POST("users/reset-password", ctr.ResetPassword)
		}
		{
			admin := frontend.Group("admin")
			admin.Use(middleware.AdminAuth())
			ctr := web.Post{}
			admin.DELETE("posts/:id", ctr.Remove)
			admin.GET("posts/form", ctr.Form)
			admin.POST("posts", ctr.Save)
		}
	}

	// api:
	api := service.Group("/api")
	api.GET("version", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"version": "0.0.1"})
	})
}
