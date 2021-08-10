package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"strings"
)

const (
	PostTypeArticle = "article" // PostTypeArticle 内容类型-文章
	PostTypePage    = "page"    // PostTypePage 内容类型-页面
	PostTypeMicro   = "micro"   // PostTypeMicro 内容类型-微博
)

type Post struct {
	ID        uint
	URL       string `gorm:"url"`
	Title     string
	Content   string
	Tags      []string `gorm:"-"`
	TagsStr   string   `gorm:"column:tags"`
	Type      string   `gorm:"column:type"`
	IsPinned  bool     `gorm:"column:is_pinned"`
	CreatedAt string   `gorm:"time"`
}

func (p *Post) AfterFind(tx *gorm.DB) (err error) {
	if p.TagsStr != "" {
		// 先尝试解析 JSON
		var tags []string
		if e := json.Unmarshal([]byte(p.TagsStr), &tags); e == nil {
			if len(tags) > 1 || tags[0] != "" {
				p.Tags = tags
			}
		} else { // 尝试使用逗号分割
			tags = strings.Split(p.TagsStr, ",")
			if len(tags) > 1 || tags[0] != "" {
				p.Tags = tags
			}
		}
	}
	return
}
