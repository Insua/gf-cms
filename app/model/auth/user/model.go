package user

import (
	"gf-cms/app/model/model"
)

type User struct {
	Name     string `gorm:"size:255;not null;uniqueIndex"`
	Body     string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	model.Model
}
