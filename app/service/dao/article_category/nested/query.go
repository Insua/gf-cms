package nested

import (
	"errors"
	"fmt"
	. "gf-cms/app/model/article/article_category"
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

	global.DB.Where(ancestorsSubSql(ac.ID)).Where("id<>(?)", ac.ID).Find(&categories)
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
