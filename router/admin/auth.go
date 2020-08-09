package admin

import (
	authController "gf-cms/app/controller/admin/auth"

	"github.com/gogf/gf/net/ghttp"
)

func auth(g *ghttp.RouterGroup) {
	g.POST("/login", authController.Login)
}
