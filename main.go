package main

import (
	"gf-cms/app/command"
	_ "gf-cms/boot"
	_ "gf-cms/router"
	"os"

	"github.com/gogf/gf/os/gcmd"

	"github.com/gogf/gf/frame/g"
)

func main() {
	if len(gcmd.GetArg(1)) > 0 {
		command.Cmd()
		os.Exit(0)
	}
	s := g.Server()
	s.Run()
}
