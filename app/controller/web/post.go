package web

import (
	"github.com/gin-gonic/gin"
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/service/post"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Post struct {
	Controller
}

func (ctr Post) Remove(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		instance.Logger().Error("无法解析id为数值", zap.String("caller", caller("post.Remove")), zap.Error(err))
		// TODO: 统一的错误显示页面
		return
	}
	if err = post.Remove(uint(idUint)); err != nil {
		// TODO: 统一的错误显示页面
		return
	}
	c.Redirect(http.StatusFound, "/")
}
