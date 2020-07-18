package router

import (
	"gf-cms/app/controller/home"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(g *ghttp.RouterGroup) {
		g.GET("/", home.Index)
	})
}
