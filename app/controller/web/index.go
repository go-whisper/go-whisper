package web

import (
	"github.com/gin-gonic/gin"
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"github.com/go-whisper/go-whisper/app/service/post"
	"go.uber.org/zap"
)

type Index struct {
	Controller
}

func (ctr Index) Index(c *gin.Context) {
	tpl := ctr.NewTemplate("index.html")
	opt := model.Option{}
	var (
		err   error
		total int64
		posts []model.Post
	)
	total, posts, err = post.List(10, 0, opt)
	if err != nil {
		instance.Logger().Error("加载首页出错", zap.String("caller", caller("Index", "Index")))
		// TODO: 输出错误页
	}
	tpl.Title = "首页 - Something"
	tpl.Data = map[string]interface{}{
		"total": total,
		"posts": posts,
	}
	ctr.Response(c, tpl)
}
