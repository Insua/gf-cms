package login

import (
	"errors"
	um "gf-cms/app/model/auth/user"
	"gf-cms/global"

	"golang.org/x/crypto/bcrypt"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

func Rule() {
	if err := gvalid.RegisterRule("login", func(value interface{}, message string, params map[string]interface{}) error {
		name := gconv.String(params["name"])
		password := gconv.String(params["password"])
		user := um.User{}
		affected := global.DB.Where(g.Map{
			"name": name,
		}).First(&user).RowsAffected
		if affected == 0 {
			return errors.New("用户不存在")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return errors.New("密码不正确")
		}
		return nil
	}); err != nil {
		panic("login rule register error")
	}
}
