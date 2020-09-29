package home

import (
	"gf-cms/app/controller/home"

	"github.com/gogf/gf/frame/g"
)

func Route() {
	s := g.Server()
	s.BindHandler("GET:/", home.Index)
}
