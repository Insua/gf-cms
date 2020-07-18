package main

import (
	_ "gf-cms/boot"
	_ "gf-cms/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
