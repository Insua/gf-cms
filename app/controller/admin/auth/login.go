package auth

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Login(r *ghttp.Request) {

}

func User(r *ghttp.Request) {
	r.Response.WriteJson(g.Map{
		"user": "user",
	})
}
