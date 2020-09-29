package home

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Index(r *ghttp.Request) {
	r.Response.WriteTpl("layout/app.html", g.Map{
		"content": "index/index.html",
	})
}
