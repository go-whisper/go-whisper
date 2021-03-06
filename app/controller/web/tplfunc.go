package web

import (
	"html/template"
	"strings"
	"time"

	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func init() {
	service := instance.WebService()
	service.SetFuncMap(template.FuncMap{
		"TimeYear":       TimeYear,
		"TimeMonth":      TimeMonth,
		"TimeDay":        TimeDay,
		"MarkdownToHTML": MarkdownToHTML,
		"StringListJoin": StringListJoin,
	})
}

func StringListJoin(strList model.StringList) string {
	return strings.Join(strList, ",")
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
	d = strings.ReplaceAll(d, model.SummarySeparator(), "")
	return template.HTML(markdown.ToHTML([]byte(d), parser, nil))
}
