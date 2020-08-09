package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gres"
)

func AdminTemplate() {
	s := g.Server()

	if gres.Contains("frontend/admin/build") {
		s.AddStaticPath("/admin", "frontend/admin/build")
		s.BindHandler("/admin/*", func(r *ghttp.Request) {
			index := gres.GetWithIndex("frontend/admin/build", []string{"index.html"})
			r.Response.Write(string(index.Content()))
		})
	}
}
