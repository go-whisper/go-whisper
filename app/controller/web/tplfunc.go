package web

import (
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"time"
)

func init() {
	service := instance.WebService()
	service.SetFuncMap(template.FuncMap{
		"TimeYear":       TimeYear,
		"TimeMonth":      TimeMonth,
		"TimeDay":        TimeDay,
		"MarkdownToHTML": MarkdownToHTML,
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

func MarkdownToHTML(d string) template.HTML {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	return template.HTML(string(markdown.ToHTML([]byte(d), parser, nil)))
}
