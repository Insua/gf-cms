package nested

import (
	"errors"
	. "gf-cms/app/model/article/article_category"
	"gf-cms/global"

	"github.com/gogf/gf/frame/g"

	"gorm.io/gorm"
)

func InsertRoot(ac *ArticleCategory) {
	type Max struct {
		Mr uint
	}
	var max Max
	global.DB.Select("MAX(Rgt) as mr").Model(ac).Scan(&max)

	ac.Lft = max.Mr + 1
	ac.Rgt = max.Mr + 2
	global.DB.Create(ac)
}

func AppendNodeToParent(p, n *ArticleCategory) error {
	if !isValidNode(p) {
		return errors.New("not a correct parent")
	}
	tx := global.DB.Begin()
	result := tx.Model(&ArticleCategory{}).Where("lft>=(?)", p.Rgt).Or("Rgt>=(?)", p.Rgt).
		Updates(g.Map{
			"lft": gorm.Expr("CASE WHEN lft>=(?) THEN lft+(2) ELSE lft END", p.Rgt),
			"Rgt": gorm.Expr("CASE WHEN Rgt>=(?) THEN Rgt+(2) ELSE lft END", p.Rgt),
		})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("update tree node error")
	}
	result = tx.Create(&ArticleCategory{
		Parent: &p.ID,
		Lft:    p.Rgt,
		Rgt:    p.Rgt + 1,
		Name:   n.Name,
	})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("insert node error")
	}
	tx.Commit()
	return nil
}

func PrependNodeToParent(p, n *ArticleCategory) error {
	if !isValidNode(p) {
		return errors.New("not a correct parent")
	}
	tx := global.DB.Begin()
	result := tx.Model(&ArticleCategory{}).Where("lft>=(?)", p.Lft+1).Or("Rgt>=(?)", p.Lft+1).
		Updates(g.Map{
			"lft": gorm.Expr("CASE WHEN lft>=(?) THEN lft+(2) ELSE lft END", p.Lft+1),
			"Rgt": gorm.Expr("CASE WHEN Rgt>=(?) THEN Rgt+(2) ELSE lft END", p.Lft+1),
		})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("update tree node error")
	}
	result = tx.Create(&ArticleCategory{
		Parent: &p.ID,
		Lft:    p.Lft + 1,
		Rgt:    p.Lft + 2,
		Name:   n.Name,
	})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("insert node error")
	}
	tx.Commit()
	return nil
}

func InsertAfterNode(neighbor, node *ArticleCategory) error {
	if !isValidNode(neighbor) {
		return errors.New("not a correct node")
	}
	tx := global.DB.Begin()
	result := tx.Model(&ArticleCategory{}).Where("lft>=(?)", neighbor.Lft+1).Or("Rgt>=(?)", neighbor.Lft+1).
		Updates(g.Map{
			"lft": gorm.Expr("CASE WHEN lft>=(?) THEN lft+(2) ELSE lft END", neighbor.Lft+1),
			"Rgt": gorm.Expr("CASE WHEN Rgt>=(?) THEN Rgt+(2) ELSE lft END", neighbor.Lft+1),
		})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("update tree node error")
	}
	result = tx.Create(&ArticleCategory{
		Parent: neighbor.Parent,
		Lft:    neighbor.Lft + 1,
		Rgt:    neighbor.Lft + 2,
		Name:   node.Name,
	})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("insert node error")
	}
	tx.Commit()
	return nil
}

func InsertBeforeNode(neighbor, node *ArticleCategory) error {
	if !isValidNode(neighbor) {
		return errors.New("not a correct node")
	}
	tx := global.DB.Begin()
	result := tx.Model(&ArticleCategory{}).Where("lft>=(?)", neighbor.Lft).Or("Rgt>=(?)", neighbor.Lft).
		Updates(g.Map{
			"lft": gorm.Expr("CASE WHEN lft>=(?) THEN lft+(2) ELSE lft END", neighbor.Lft),
			"Rgt": gorm.Expr("CASE WHEN Rgt>=(?) THEN Rgt+(2) ELSE lft END", neighbor.Lft),
		})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("update tree node error")
	}
	result = tx.Create(&ArticleCategory{
		Parent: neighbor.Parent,
		Lft:    neighbor.Lft,
		Rgt:    neighbor.Lft + 1,
		Name:   node.Name,
	})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("insert node error")
	}
	tx.Commit()
	return nil
}
