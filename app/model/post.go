package model

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

const (
	PostTypeArticle = "article" // PostTypeArticle 内容类型-文章
	PostTypePage    = "page"    // PostTypePage 内容类型-页面
	PostTypeMicro   = "micro"   // PostTypeMicro 内容类型-微博
)

type Post struct {
	ID           uint
	URL          string `gorm:"url"`
	Title        string
	Content      string
	Tags         []string  `gorm:"-"`
	TagsStr      string    `gorm:"column:tags"`
	CreatedAt    time.Time `gorm:"-"`
	CreatedAtInt int64     `gorm:"column:created_at"`
	Type         string    `gorm:"column:type"`
	IsPinned     bool      `gorm:"column:is_pinned"`
}

func (p *Post) AfterFind(tx *gorm.DB) (err error) {
	if p.TagsStr != "" {
		p.Tags = strings.Split(p.TagsStr, ",")
	}
	if p.CreatedAtInt > 0 {
		p.CreatedAt = time.Unix(p.CreatedAtInt, 0)
	}
	return
}
