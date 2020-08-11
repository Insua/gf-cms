package auth

import (
	"gf-cms/app/model/auth/users"

	"github.com/gogf/gf/net/ghttp"
)

func Login(r *ghttp.Request) {

}

func User(r *ghttp.Request) {
	isLogin := false
	if id := r.Session.GetInt("user-id"); id > 0 {
		count, err := users.FindCount("id=(?)", id)
		if err == nil && count > 0 {
			isLogin = true
		}
	}
	type Response struct {
		User *int `json:"user"`
	}
	var response Response
	if isLogin {
		user := r.Session.GetInt("user-id")
		response.User = &user
	}
	r.Response.WriteJson(response)
}

func Test(r *ghttp.Request) {
	user, _ := users.FindAll()
	r.Response.WriteJson(user)
}
