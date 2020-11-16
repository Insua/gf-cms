package nested

import (
	"errors"
	"fmt"
	. "gf-cms/app/model/article/article_category"
	"gf-cms/app/service/orm"
	"gf-cms/global"

	"gorm.io/gorm"
)

func ancestorsSubSql(id uint) string {
	subStmt := global.DB.Session(&gorm.Session{DryRun: true}).Model(&ArticleCategory{}).
		Select("rgt").Where(fmt.Sprintf("id=(%v)", id)).First(&ArticleCategory{}).Statement
	return fmt.Sprintf("(%s) between lft AND rgt", subStmt.SQL.String())
}

func AncestorsOf(ac *ArticleCategory) ([]*ArticleCategory, error) {
	if !isValidNode(ac) {
		return nil, errors.New("not a correct node")
	}
	categories := make([]*ArticleCategory, 0)

	global.DB.Model("").Where(ancestorsSubSql(ac.ID)).Where("id<>(?)", ac.ID).Find(&categories)
	return categories, nil
}

func AncestorsAndSelf(ac *ArticleCategory) ([]*ArticleCategory, error) {
	if !isValidNode(ac) {
		return nil, errors.New("not a correct node")
	}
	categories := make([]*ArticleCategory, 0)

	global.DB.Where(ancestorsSubSql(ac.ID)).Find(&categories)
	return categories, nil
}

func DescendantsOf(ac *ArticleCategory) ([]*ArticleCategory, error) {
	if !isValidNode(ac) {
		return nil, errors.New("not a correct node")
	}
	categories := make([]*ArticleCategory, 0)

	global.DB.Where("lft between (?) AND (?)", ac.Lft+1, ac.Rgt).Find(&categories)
	return categories, nil
}

func DescendantsAndSelf(ac *ArticleCategory) ([]*ArticleCategory, error) {
	if !isValidNode(ac) {
		return nil, errors.New("not a correct node")
	}
	categories := make([]*ArticleCategory, 0)

	global.DB.Where("lft between (?) AND (?)", ac.Lft, ac.Rgt).Find(&categories)
	return categories, nil
}

func NextSibling(ac *ArticleCategory) (*ArticleCategory, error) {
	if !isValidNode(ac) {
		return nil, errors.New("not a correct node")
	}

	var sibling ArticleCategory
	if result := global.DB.Scopes(orm.IsNullOrEqual("parent", ac.Parent)).
		Where("lft>(?)", ac.Lft).Order("lft ASC").First(&sibling); result.RowsAffected == 0 {
		return nil, nil
	}

	return &sibling, nil
}

func PrevSibling(ac *ArticleCategory) (*ArticleCategory, error) {
	if !isValidNode(ac) {
		return nil, errors.New("not a correct node")
	}

	var sibling ArticleCategory
	if result := global.DB.Scopes(orm.IsNullOrEqual("parent", ac.Parent)).
		Where("lft<(?)", ac.Lft).Order("lft DESC").First(&sibling); result.RowsAffected == 0 {
		return nil, nil
	}

	return &sibling, nil
}
