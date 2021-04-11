package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-whisper/go-whisper/app/controller/web"
)

func InitUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if name, ok := c.Cookie(web.UserCookieNamePrefix + "name"); ok == nil && name != "" {
			c.Set("user_name", name)
		}
		c.Next()
	}
}
