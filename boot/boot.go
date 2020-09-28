package boot

import (
	"gf-cms/global"
	_ "gf-cms/packed"
	"os"

	"github.com/gookit/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gogf/gf/frame/g"
)

func init() {
	db, err := gorm.Open(mysql.Open(g.Cfg().GetString("gorm.mysql.dsn")), &gorm.Config{})
	if err != nil {
		color.Red.Println(err)
		os.Exit(0)
	} else {
		if sqlDB, err := db.DB(); err == nil {
			if err := sqlDB.Ping(); err != nil {
				color.Red.Println(err)
				os.Exit(0)
			}
		} else {
			color.Red.Println(err)
			os.Exit(0)
		}
		if g.Cfg().GetBool("gorm.mysql.debug") {
			global.DB = db.Debug()
		} else {
			global.DB = db
		}
	}
}
