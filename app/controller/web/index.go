package web

import (
	"github.com/gin-gonic/gin"
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"github.com/go-whisper/go-whisper/app/service/post"
	"go.uber.org/zap"
	"html/template"
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
	p, _ := ctr.GetQueryInt(c, "page", 1)
	pageSize := model.GetSite().PageSize
	total, posts, err = post.List(pageSize, pageSize*p-pageSize, opt)
	if err != nil {
		instance.Logger().Error("加载首页出错", zap.String("caller", caller("Index", "Index")))
		// TODO: 输出错误页
	}

	pagination := NewPagination(c.Request, int(total), pageSize)

	tpl.Title = "首页 - " + tpl.Site.Name
	tpl.Data = gin.H{
		"total": total * 10,
		"posts": posts,
		"p":     p,
		"page":  template.HTML(pagination.Pages()),
	}
	ctr.Response(c, tpl)
}
