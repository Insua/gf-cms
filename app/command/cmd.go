package command

import (
	"gf-cms/app/command/migrate"

	"github.com/gogf/gf/os/gcmd"
	"github.com/gookit/color"
)

func Cmd() {
	command := gcmd.GetArg(1)

	switch command {
	case "migrate":
		sub := gcmd.GetArg(2)
		switch sub {
		case "new":
			fileName := gcmd.GetArg(3)
			if len(fileName) == 0 {
				color.Error.Prompt("Doesn't have fileName")
			} else {
				migrate.New(fileName)
			}
		case "up":
			migrate.Up()
		case "down":
			migrate.Down()
		default:
			color.Warn.Prompt("Wrong Process method")
		}

	default:
		color.Warn.Prompt("Not correct command")
	}
}
