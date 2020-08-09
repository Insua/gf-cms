package home

import (
	"github.com/gogf/gf/net/ghttp"
)

func Index(r *ghttp.Request) {
	r.Response.Write("index")
}
