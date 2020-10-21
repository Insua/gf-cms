package migrations

import (
	"gf-cms/global"
	"time"
)

type Migration202010171704188781CreateArticleCategoriesTable struct {
	Name string
}

func init() {
	m := Migration202010171704188781CreateArticleCategoriesTable{
		Name: "202010171704188781_create_article_categories_table",
	}
	Migrations = append(Migrations, &m)
}

type ArticleCategory struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Lft       *uint  `gorm:"default:0;not null"`
	Rft       *uint  `gorm:"default:0;not null"`
	Parent    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Migration202010171704188781CreateArticleCategoriesTable) Up() error {
	return global.DB.Migrator().CreateTable(&ArticleCategory{})
}

func (m *Migration202010171704188781CreateArticleCategoriesTable) Down() error {
	return global.DB.Migrator().DropTable(&ArticleCategory{})
}
