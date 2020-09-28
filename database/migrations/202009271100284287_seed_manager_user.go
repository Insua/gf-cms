package migrations

import (
	"errors"
	um "gf-cms/app/model/auth/user"
	"gf-cms/global"

	"golang.org/x/crypto/bcrypt"
)

type Migration202009271100284287SeedManagerUser struct {
	Name string
}

func init() {
	m := Migration202009271100284287SeedManagerUser{
		Name: "202009271100284287_seed_manager_user",
	}
	Migrations = append(Migrations, &m)
}

func (m *Migration202009271100284287SeedManagerUser) Up() error {
	hash, err := bcrypt.GenerateFromPassword([]byte("gfcms"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := um.User{
		Name:     "manager",
		Body:     "管理员",
		Password: string(hash),
	}
	if affected := global.DB.Create(&user).RowsAffected; affected == 0 {
		return errors.New("创建管理员失败")
	}
	return nil
}

func (m *Migration202009271100284287SeedManagerUser) Down() error {
	if global.DB.Where("name = ?", "manager").Delete(&um.User{}).RowsAffected == 0 {
		return errors.New("删除管理员失败")
	}
	return nil
}
