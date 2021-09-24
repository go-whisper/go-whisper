package model

import (
	"reflect"
	"strconv"
	"time"

	"github.com/go-whisper/go-whisper/app/instance"
	"go.uber.org/zap"
)

type SiteParameter struct {
	Name        string `db-option:"name"`
	Description string `db-option:"description"`
	Domain      string `db-option:"domain"`
	PageSize    int    `db-option:"page_size"`
	Separator   string `db-option:"separator"`
}

type DBSiteParameter struct {
	ID     uint
	Option string
	Value  string
}

func (DBSiteParameter) TableName() string {
	return "site_parameters"
}

func GetSiteParameter() *SiteParameter {
	k := "site"
	c := instance.Cache()
	if c.Exists(k) {
		res, e := c.Value(k)
		if e == nil {
			if v, ok := res.Data().(*SiteParameter); ok {
				return v
			}
		}
	}
	site := SiteParameter{}
	var parameters []DBSiteParameter
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

func UpdateSiteParameter(sp SiteParameter) bool {
	fields := reflect.TypeOf(sp)
	values := reflect.ValueOf(&sp)
	valuesElem := values.Elem()
	var opt string
	var vStr string
	for i := 0; i < fields.NumField(); i++ {
		opt = fields.Field(i).Tag.Get("db-option")
		v := valuesElem.FieldByName(fields.Field(i).Name)
		switch v.Kind() {
		case reflect.String:
			vStr = v.String()
		case reflect.Int:
			vStr = strconv.Itoa(int(v.Int()))
		}
		if err := instance.DB().Model(&sp).Where("option=?", opt).UpdateColumn("value", vStr).Error; err != nil {
			instance.Logger().Error("UpdateSiteParameter fail", zap.String("option", opt), zap.Error(err))
			return false
		}
	}
	UpdateLastChange(instance.DB())
	return true
}
