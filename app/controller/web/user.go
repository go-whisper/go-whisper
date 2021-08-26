package web

import (
	"github.com/gin-gonic/gin"
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/service/user"
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
		instance.Logger().Warn("用户登录失败", zap.String("name", c.PostForm("name")))
		ctr.Response(c, tpl)
		return
	}
	age := 24 * 3600
	if c.PostForm("remember") == "yes" {
		age *= 30
	}
	c.SetCookie(UserCookieNamePrefix+"name", u.Name, age, "/", "", false, true)
	c.SetCookie(UserCookieNamePrefix+"id", strconv.Itoa(int(u.ID)), age, "/", "", false, true)
	instance.Logger().Info("用户登录成功", zap.String("name", c.PostForm("name")))
	c.Redirect(http.StatusFound, "/")
}

func (ctr User) Logout(c *gin.Context) {
	c.SetCookie(UserCookieNamePrefix+"name", "", -1, "/", "", false, true)
	c.SetCookie(UserCookieNamePrefix+"id", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/")
}

func (ctr User) ResetPasswordForm(c *gin.Context) {
	tpl := ctr.NewTemplate("reset-password.html")
	tpl.Title = "设置密码" + tpl.Site.Name
	tpl.Data = gin.H{}
	ctr.Response(c, tpl)
}
func (ctr User) ResetPassword(c *gin.Context) {
	if len(c.PostForm("password")) < 6 {
		ctr.Error(c, "新密码不能少于 6 个字符")
		return
	}
	name, _ := c.Cookie(UserCookieNamePrefix + "name")
	if err := user.UpdatePassword(name, c.PostForm("original_password"), c.PostForm("password")); err != nil {
		ctr.Error(c, err.Error())
		return
	}
	ctr.Success(c, "密码修改成功")
}
