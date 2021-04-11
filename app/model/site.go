package model

import (
	"reflect"
	"strconv"
	"time"

	"github.com/go-whisper/go-whisper/app/instance"
	"go.uber.org/zap"
)

type Site struct {
	Name        string `db-option:"name"`
	Description string `db-option:"description"`
	Domain      string `db-option:"domain"`
	PageSize    int    `db-option:"page_size"`
	Separator   string `db-option:"separator"`
}

type SiteParameter struct {
	ID     uint
	Option string
	Value  string
}

func (SiteParameter) TableName() string {
	return "site_parameters"
}

func GetSite() *Site {
	k := "site"
	c := instance.Cache()
	if c.Exists(k) {
		res, e := c.Value(k)
		if e == nil {
			if v, ok := res.Data().(*Site); ok {
				return v
			}
		}
	}
	site := Site{}
	var parameters []SiteParameter
	if err := instance.DB().Find(&parameters).Error; err != nil {
		instance.Logger().Error("读取站点参数失败", zap.Error(err))
		return &site
	}
	if len(parameters) < 3 {
		instance.Logger().Error("站点参数可用项过少", zap.Int("len", len(parameters)))
		return &site
	}

	fields := reflect.TypeOf(site)
	values := reflect.ValueOf(&site)
	valuesElem := values.Elem()
	var opt string
	for i := 0; i < fields.NumField(); i++ {
		opt = fields.Field(i).Tag.Get("db-option")
		for _, p := range parameters {
			if p.Option == opt {
				v := valuesElem.FieldByName(fields.Field(i).Name)
				switch v.Kind() {
				case reflect.String:
					v.SetString(p.Value)
				case reflect.Int:
					i64, e := strconv.ParseInt(p.Value, 10, 64)
					if e != nil {
						instance.Logger().Error("站点参数配置错误,无法转换成数值型", zap.String("value", p.Value), zap.Error(e))
					}
					v.SetInt(i64)
				}
			}
		}
	}
	c.Add(k, time.Minute*10, &site)
	return &site
}
