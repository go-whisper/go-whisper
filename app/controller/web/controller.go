package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Controller struct {
	//
}

type Template struct {
	Path        string
	Title       string
	Keywords    string
	Description string
	Author      string
	Data        map[string]interface{}
}

func (ctr Controller) NewTemplate(path string) *Template {
	tplName := "default/" // 暂不支持多模板
	if path != "" && !strings.HasPrefix(path, tplName) {
		path = tplName + path
	}
	if path != "" && !strings.HasSuffix(path, ".html") {
		path += ".html"
	}
	return &Template{Path: path}
}

func (ctr Controller) Response(c *gin.Context, tpl *Template) {
	c.HTML(http.StatusOK, tpl.Path, *tpl)
}
