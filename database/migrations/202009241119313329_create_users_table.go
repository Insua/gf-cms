package migrations

import (
	"gf-cms/global"
	"time"
)

type Migration202009241119313329CreateUsersTable struct {
	Name string
}

func init() {
	m := Migration202009241119313329CreateUsersTable{
		Name: "202009241119313329_create_users_table",
	}
	Migrations = append(Migrations, &m)
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null;uniqueIndex"`
	Body      string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Migration202009241119313329CreateUsersTable) Up() error {
	return global.DB.Migrator().CreateTable(&User{})
}

func (m *Migration202009241119313329CreateUsersTable) Down() error {
	return global.DB.Migrator().DropTable(&User{})
}
