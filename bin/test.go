package main

import (
	article_category2 "gf-cms/app/model/article/article_category"
	"gf-cms/app/service/dao/article_category"
	_ "gf-cms/boot"
)

func main() {
	ac := &article_category2.ArticleCategory{
		Name: "test",
	}
	article_category.InsertRoot(ac)
}
