package article_category

import "gf-cms/app/model/model"

type ArticleCategory struct {
	model.Model
	Name   string `gorm:"size:255"`
	Lft    *uint  `gorm:"default:0;not null"`
	Rft    *uint  `gorm:"default:0;not null"`
	Parent uint
}
