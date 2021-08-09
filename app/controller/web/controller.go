package web

import (
	"github.com/go-whisper/go-whisper/app/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const UserCookieNamePrefix = "user."

type Controller struct {
	//
}

type Template struct {
	Site        *model.Site
	Path        string // 要使用的模板路径
	Title       string // 页面 title 标签
	Keywords    string // 页面 meta 标签: Keywords
	Description string // 页面 meta 标签: Description
	Author      string // 页面 meta 标签: Author
	IsLoggedIn  bool   // 当前用户是否已登录
	Data        gin.H
}

// NewTemplate 返回一个 Template
// `path` 不需要包含 模板主题名
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
	if v, has := c.Get("user_name"); has {
		if s, ok := v.(string); ok && s != "" {
			tpl.IsLoggedIn = true
		}
	}
	c.HTML(http.StatusOK, tpl.Path, *tpl)
}

// GetQueryInt 返回查询参数中 key 对应的值
// 如果找不到 key 或者 key 的值不是有效数值，则返回的第2个参数返回 false
// 可以指定 defVal 参数作为默认值返回
func (*Controller) GetQueryInt(c *gin.Context, key string, defVal ...int) (int, bool) {
	def := 0
	if len(defVal) == 1 {
		def = defVal[0]
	}
	val, has := c.GetQuery(key)
	if !has {
		return def, false
	}
	if v, err := strconv.Atoi(val); err == nil {
		return v, true
	}
	return def, false
}

// GetParamInt 从路由参数中返回 key 对应的值
// 规则与 GetQueryInt() 函数相同
func (*Controller) GetParamInt(c *gin.Context, key string, defVal ...int) (int, bool) {
	def := 0
	if len(defVal) == 1 {
		def = defVal[0]
	}
	strVal := c.Param(key)
	if v, err := strconv.Atoi(strVal); err == nil {
		return v, true
	}
	return def, false
}
