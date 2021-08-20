package web

import (
	"github.com/gin-gonic/gin"
	"github.com/go-whisper/go-whisper/app/model"
)

type SiteParameter struct {
	Controller
}

func (ctr SiteParameter) Form(c *gin.Context) {
	tpl := ctr.NewTemplate("site-parameter.html")
	tpl.Title = "参数设置 " + tpl.Site.Name
	tpl.Data = gin.H{"parameters": model.GetSiteParameter()}
	ctr.Response(c, tpl)
}

func (ctr SiteParameter) Save(c *gin.Context) {
	req := siteparameterRequest{}
	if err := c.ShouldBind(&req); err != nil {
		ctr.Error(c, "参数错误:"+err.Error())
		return
	}
	parameter := model.SiteParameter{
		Name:        req.Name,
		Domain:      req.Domain,
		Description: req.Description,
		PageSize:    req.PageSize,
		Separator:   req.Separator,
	}
	if !model.UpdateSiteParameter(parameter) {
		ctr.Error(c, "操作失败")
		return
	}
	ctr.Success(c, "操作成功")
}
