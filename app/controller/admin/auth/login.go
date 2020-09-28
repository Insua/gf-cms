package auth

import (
	um "gf-cms/app/model/auth/user"
	"gf-cms/app/request/admin/auth"
	"gf-cms/app/service/http"
	"gf-cms/global"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/util/gvalid"

	"github.com/gogf/gf/net/ghttp"
)

func Login(r *ghttp.Request) {
	req := new(auth.LoginRequest)
	errParse := r.Parse(req)
	if errParse != nil {
		http.Error422(r, []string{"参数错误"})
	}

	errValid := gvalid.CheckStruct(req, auth.LoginRequestRules, auth.LoginRequestMessages)
	if errValid != nil {
		http.ErrorValid(r, errValid)
	}

	user := um.User{}
	global.DB.Where(g.Map{
		"name": req.Name,
	}).First(&user)
	errSession := r.Session.Set("user-id", user.ID)
	if errSession != nil {
		http.Error500(r)
	}

	r.Response.WriteJsonExit(g.Map{
		"id":   user.ID,
		"name": user.Name,
		"body": user.Body,
	})
}

func User(r *ghttp.Request) {
	isLogin := false
	if id := r.Session.GetInt("user-id"); id > 0 {
		var count int64
		if global.DB.Model(&um.User{}).Where("id = ?", id).Count(&count); count > 0 {
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
