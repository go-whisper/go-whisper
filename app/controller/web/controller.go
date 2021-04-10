package web

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	//
}

type Blog struct {
	Name        string
	Description string
	Domain      string
}

type Template struct {
	Blog        Blog
	Path        string
	Title       string
	Keywords    string
	Description string
	Author      string
	Data        map[string]interface{}
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
	tpl.Blog = Blog{Name: "Blog Name", Description: "Blog Description", Domain: "https://domain.com"}
	return tpl
}

func (ctr Controller) Response(c *gin.Context, tpl *Template) {
	c.HTML(http.StatusOK, tpl.Path, *tpl)
}
