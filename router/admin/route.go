package admin

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Route() {
	s := g.Server()
	s.Group("/api/admin", func(g *ghttp.RouterGroup) {
		g.Group("/auth", func(g *ghttp.RouterGroup) {
			auth(g)
		})
	})
}
