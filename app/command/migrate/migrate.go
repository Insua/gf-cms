package migrate

import (
	_ "gf-cms/boot"
	ms "gf-cms/database/migrations"
	"gf-cms/global"

	migrate "github.com/Insua/gorm-migrate"
	"github.com/gookit/color"
)

func Up() {
	migrations := ms.Migrations
	err := migrate.Up(global.DB, migrations)
	if err != nil {
		color.Red.Println(err)
	}
}

func Down() {
	migrations := ms.Migrations
	err := migrate.Down(global.DB, migrations)
	if err != nil {
		color.Red.Println(err)
	}
}
