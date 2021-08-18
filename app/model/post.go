package model

import (
	"strings"

	"gorm.io/gorm"
)

const (
	PostTypeArticle = "article" // PostTypeArticle 内容类型-文章
	PostTypePage    = "page"    // PostTypePage 内容类型-页面
	PostTypeMicro   = "micro"   // PostTypeMicro 内容类型-微博
)

type Post struct {
	ID      uint
	URL     string `gorm:"url"`
	Title   string
	Content string
	Summary string `gorm:"-"`
	Tags    StringList
	// TagsStr   string   `gorm:"column:tags"`
	Type      string `gorm:"column:type"`
	IsPinned  bool   `gorm:"column:is_pinned"`
	CreatedAt string `gorm:"time"`
}

func (p *Post) AfterFind(tx *gorm.DB) (err error) {
	p.Summary = p.GetSummary()
	return
}

func (p Post) GetSummary() string {
	tmp := strings.Split(p.Content, SummarySeparator())
	return tmp[0]
}

var summarySeparator string

func SummarySeparator() string {
	if summarySeparator != "" {
		return summarySeparator
	}
	site := GetSite()
	summarySeparator = site.Separator
	return summarySeparator
}
