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
	Pages   int    `gorm:"-"` // 内容分页的总页数
	Tags    StringList
	// TagsStr   string   `gorm:"column:tags"`
	Type      string `gorm:"column:type"`
	IsPinned  bool   `gorm:"column:is_pinned"`
	CreatedAt string `gorm:"time"`
}

func (p *Post) AfterFind(*gorm.DB) (err error) {
	tmp := strings.Split(p.Content, SummarySeparator())
	p.Summary = tmp[0]
	p.Pages = len(tmp)
	return
}

var summarySeparator string

func SummarySeparator() string {
	if summarySeparator != "" {
		return summarySeparator
	}
	site := GetSiteParameter()
	summarySeparator = site.Separator
	return summarySeparator
}
