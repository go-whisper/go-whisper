package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-whisper/go-whisper/app/controller/web"
	"net/http"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		stop := func() {
			c.Redirect(http.StatusTemporaryRedirect, "/users/login")
			//ctr.DisplayLoginForm(c, "请登录")
			c.Abort()
		}
		// TODO: 验证 cookie
		if _, err := c.Cookie(web.UserCookieNamePrefix + "name"); err == nil {
			c.Next()
		} else {
			stop()
		}
	}
}
