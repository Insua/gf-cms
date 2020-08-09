package migrate

import (
	"database/sql"
	"fmt"
	"gf-cms/library/util/slice"
	"os"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gookit/color"
	migrate "github.com/rubenv/sql-migrate"
)

func dbMigrate(direction string) {
	if !slice.InString([]string{"up", "down"}, direction) {
		color.Error.Prompt("Wrong migrate process")
		os.Exit(1)
	}
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
	dir := migrate.Up
	max := 0
	if direction == "down" {
		dir = migrate.Down
		max = 1
	}
	n, errMigrate := migSet.ExecMax(db, "mysql", migrations, dir, max)
	if errMigrate != nil {
		color.Error.Prompt(errMigrate.Error())
		os.Exit(1)
	}
	if n > 0 {
		color.Info.Prompt(fmt.Sprintf("Migrate %s success %d", direction, n))
	} else {
		color.Warn.Prompt(fmt.Sprintf("Nothing to migrate %s", direction))
	}
}
