package router

import (
	"gf-cms/router/admin"
	"gf-cms/router/home"

	"github.com/gogf/gf/frame/g"
)

func init() {
	s := g.Server()
	s.SetSessionIdName(g.Cfg().GetString("server.SessionIdName", "gf-cms-session"))
	AdminTemplate()
	home.Route()
	admin.Route()
}
