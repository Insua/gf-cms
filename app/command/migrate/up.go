package migrate

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"

	"github.com/gogf/gf/frame/g"

	_ "github.com/go-sql-driver/mysql"

	migrate "github.com/rubenv/sql-migrate"
)

func Up() {
	migrations := &migrate.FileMigrationSource{
		Dir: "database/migrations",
	}
	link := g.Cfg().GetString("database.link")
	conn := fmt.Sprintf("%s?parseTime=true&loc=Local", link[strings.Index(link, ":")+1:])
	db, err := sql.Open("mysql", conn)
	if err != nil {
		color.Error.Prompt("Connect Error")
		os.Exit(1)
	}
	var migSet = migrate.MigrationSet{
		TableName: "migrations",
	}
	n, errMigrate := migSet.Exec(db, "mysql", migrations, migrate.Up)
	if errMigrate != nil {
		color.Error.Prompt(errMigrate.Error())
		os.Exit(1)
	}
	if n > 0 {
		color.Info.Prompt(fmt.Sprintf("Migrate up success %d", n))
	} else {
		color.Warn.Prompt("Noting to migrate up")
	}
}
