package web

import (
	"net/http"
	"strings"

	"github.com/go-whisper/go-whisper/app/model"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	//
}

type Template struct {
	Site        *model.Site
	Path        string
	Title       string
	Keywords    string
	Description string
	Author      string
	Data        gin.H
}

func (ctr Controller) NewTemplate(path string) *Template {
	tplName := "crude/" // 暂不支持多模板
	if path != "" && !strings.HasPrefix(path, tplName) {
		path = tplName + path
	}
	if path != "" && !strings.HasSuffix(path, ".html") {
		path += ".html"
	}
	tpl := &Template{Path: path}
	tpl.Site = model.GetSite()
	return tpl
}

func (ctr Controller) Response(c *gin.Context, tpl *Template) {
	c.HTML(http.StatusOK, tpl.Path, *tpl)
}
