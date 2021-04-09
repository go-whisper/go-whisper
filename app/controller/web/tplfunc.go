package web

import (
	"github.com/go-whisper/go-whisper/app/instance"
	"html/template"
	"time"
)

func init() {
	service := instance.WebService()
	service.SetFuncMap(template.FuncMap{
		"TimeYear":  TimeYear,
		"TimeMonth": TimeMonth,
		"TimeDay":   TimeDay,
	})
}

func TimeYear(t time.Time) int {
	return t.Year()
}

func TimeMonth(t time.Time) string {
	return t.Month().String()
}

func TimeDay(t time.Time) int {
	return t.Day()
}
