package migrate

import (
	migrate "github.com/Insua/gorm-migrate"
	"github.com/gookit/color"
)

func New(fileName string) {
	err := migrate.Create("database/migrations", "migrations", fileName)
	if err != nil {
		color.Red.Println(err)
	} else {
		color.Green.Println(fileName + " has be created")
	}
}
