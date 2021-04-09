package model

import (
	"strings"
	"time"

	"gorm.io/gorm"
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
