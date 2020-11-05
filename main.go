package main

import (
	_ "gf-cms/boot"
	_ "gf-cms/router"

	"github.com/gogf/gf/os/gres"

	"github.com/gogf/gf/frame/g"
)

func main() {
	gres.Dump()
	s := g.Server()
	s.Run()
}
