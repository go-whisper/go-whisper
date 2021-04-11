package web

import (
	"github.com/gin-gonic/gin"
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/service/user"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type User struct {
	Controller
}

func (ctr User) LoginForm(c *gin.Context) {
	tpl := ctr.NewTemplate("login.html")
	tpl.Title = "登录 - " + tpl.Site.Name
	ctr.Response(c, tpl)
}

func (ctr User) Login(c *gin.Context) {
	u, err := user.Login(c.PostForm("name"), c.PostForm("password"))
	if err != nil {
		tpl := ctr.NewTemplate("login.html")
		tpl.Title = "登录 - " + tpl.Site.Name
		tpl.Data = gin.H{"message": err.Error()}
		ctr.Response(c, tpl)
		return
	}
	age := 24 * 3600
	if c.PostForm("remember") == "yes" {
		age *= 30
	}
	c.SetCookie(UserCookieNamePrefix+"name", u.Name, age, "/", "", false, false)
	c.SetCookie(UserCookieNamePrefix+"id", strconv.Itoa(int(u.ID)), age, "/", "", false, false)
	instance.Logger().Info("用户登录成功", zap.String("name", c.PostForm("name")))
	c.Redirect(http.StatusFound, "/")
}

func (ctr User) ResetPassword(c *gin.Context) {
	resetToken := viper.GetString("secure.resetToken")
	if len(resetToken) != 128 {
		c.Data(http.StatusUnprocessableEntity, "text/html", []byte("请检查配置文件中 [secure.resetToken] 必须是长度为128的字符串"))
		return
	}
	if c.PostForm("reset_token") != resetToken {
		c.Data(http.StatusUnprocessableEntity, "text/html", []byte("Token 无效"))
		return
	}
	if len(c.PostForm("name")) < 2 {
		c.Data(http.StatusUnprocessableEntity, "text/html", []byte("name 参数不能少于 2 个字符"))
		return
	}
	if len(c.PostForm("password")) < 6 {
		c.Data(http.StatusUnprocessableEntity, "text/html", []byte("新密码不能少于 6 个字符"))
		return
	}
	if err := user.UpdatePassword(c.PostForm("name"), c.PostForm("password")); err != nil {
		c.Data(http.StatusInternalServerError, "text/html", []byte(err.Error()))
		return
	}
	c.Data(http.StatusOK, "text/html", []byte("操作成功"))
}
