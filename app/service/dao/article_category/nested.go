package article_category

import (
	"errors"
	"gf-cms/app/model/article/article_category"
	"gf-cms/global"

	"github.com/gogf/gf/frame/g"

	"gorm.io/gorm"
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

func AppendNodeToParent(p, n *article_category.ArticleCategory) error {
	if p.ID == 0 || p.Lft == 0 || p.Rft == 0 || p.Lft >= p.Rft {
		return errors.New("not a correct parent")
	}
	tx := global.DB.Begin()
	result := tx.Model(&article_category.ArticleCategory{}).Where("lft>=(?)", p.Rft).Or("rft>=(?)", p.Rft).
		Updates(g.Map{
			"lft": gorm.Expr("CASE WHEN lft>=(?) THEN lft+(2) ELSE lft END", p.Rft),
			"rft": gorm.Expr("CASE WHEN rft>=(?) THEN rft+(2) ELSE lft END", p.Rft),
		})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("update tree node error")
	}
	result = tx.Create(&article_category.ArticleCategory{
		Parent: &p.ID,
		Lft:    p.Rft,
		Rft:    p.Rft + 1,
		Name:   n.Name,
	})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("insert node error")
	}
	tx.Commit()
	return nil
}

func PrependNodeToParent(p, n *article_category.ArticleCategory) error {
	if p.ID == 0 || p.Lft == 0 || p.Rft == 0 || p.Lft >= p.Rft {
		return errors.New("not a correct parent")
	}
	tx := global.DB.Begin()
	result := tx.Model(&article_category.ArticleCategory{}).Where("lft>=(?)", p.Lft+1).Or("rft>=(?)", p.Lft+1).
		Updates(g.Map{
			"lft": gorm.Expr("CASE WHEN lft>=(?) THEN lft+(2) ELSE lft END", p.Lft+1),
			"rft": gorm.Expr("CASE WHEN rft>=(?) THEN rft+(2) ELSE lft END", p.Lft+1),
		})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("update tree node error")
	}
	result = tx.Create(&article_category.ArticleCategory{
		Parent: &p.ID,
		Lft:    p.Lft + 1,
		Rft:    p.Lft + 2,
		Name:   n.Name,
	})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("insert node error")
	}
	tx.Commit()
	return nil
}
