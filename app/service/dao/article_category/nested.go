package article_category

import (
	"gf-cms/app/model/article/article_category"
	"gf-cms/global"
)

func InsertRoot(ac *article_category.ArticleCategory) {
	type Max struct {
		Mr uint
	}
	var max Max
	global.DB.Select("MAX(rft) as mr").Model(ac).Scan(&max)

	ac.Lft = max.Mr + 1
	ac.Rft = max.Mr + 2
	global.DB.Create(ac)
}
