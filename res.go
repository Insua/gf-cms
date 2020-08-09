package main

import (
	"fmt"
	_ "gf-cms/boot"

	"github.com/gogf/gf/os/gres"
)

func main() {
	gres.Dump()
	fs := gres.ScanDir("frontend/admin/build", "*", true)
	for _, v := range fs {
		fmt.Println(v.Name())
	}
}
