package web

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"github.com/go-whisper/go-whisper/app/service/post"
	"go.uber.org/zap"
)

type Post struct {
	Controller
}

func (ctr Post) Detail(c *gin.Context) {
	var (
		p   model.Post
		err error
	)
	id, _ := ctr.GetParamInt(c, "flag")
	if id > 0 {
		p, err = post.Detail(uint(id))
	} else {
		p, err = post.DetailByURL(c.Param("flag"))
	}
	if err != nil {
		ctr.Error(c, "读取数据出错，请稍候再试")
		return
	}
	ctr.detail(c, p, "")
}

func (ctr Post) DetailForPage(c *gin.Context) {
	p, err := post.DetailByURL(c.Param("page"))
	if err != nil {
		ctr.Error(c, "读取数据出错，请稍候再试")
		return
	}
	ctr.detail(c, p, "page")
}

func (ctr Post) detail(c *gin.Context, p model.Post, contentType string) {
	tpl := ctr.NewTemplate("post-detail.html")
	tpl.Title = p.Title
	tpl.Data = gin.H{
		"post":        p,
		"id":          p.ID,
		"contentType": contentType,
	}
	ctr.Response(c, tpl)
}

func (ctr Post) Form(c *gin.Context) {
	id, _ := ctr.GetQueryInt(c, "id", 0)
	p, _ := post.Detail(uint(id))
	tpl := ctr.NewTemplate("post-form.html")
	tpl.Title = "编辑内容 - " + tpl.Site.Name
	tpl.Data = gin.H{
		"post":             p,
		"id":               id,
		"summarySeparator": model.SummarySeparator(),
	}
	ctr.Response(c, tpl)
}

func (ctr Post) Save(c *gin.Context) {
	req := postRequest{}
	if err := c.ShouldBind(&req); err != nil {
		ctr.Error(c, "参数错误:"+err.Error())
		return
	}
	req.Content = strings.ReplaceAll(req.Content, "\r\n", "\n")
	p := model.Post{
		Title:    req.Title,
		Content:  req.Content,
		Tags:     model.NewStringList(req.Tags),
		IsPinned: req.IsPinned,
		URL:      req.URL,
	}
	id, _ := ctr.GetQueryInt(c, "id", 0)
	var err error
	if id == 0 {
		err = post.Create(&p)
	} else {
		err = post.Update(uint(id), &p)
	}
	if err != nil {
		ctr.Error(c, "处理数据出错:"+err.Error())
		return
	}
	ctr.Success(c, "内容已保存。")
}

func (ctr Post) Remove(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		instance.Logger().Error("无法解析id为数值", zap.String("caller", caller("post.Remove")), zap.Error(err))
		ctr.JsonFail(c, "无法获取ID参数", 400)
		return
	}
	if err = post.Remove(uint(idUint)); err != nil {
		ctr.JsonFail(c, "操作错误", 500)
		return
	}
	ctr.JsonSuccess(c, gin.H{"id": id})
}
