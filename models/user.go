package models

import (
	"app/dao"
	"app/pkg/e"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"column:name" binding:"required"`
	Password string `gorm:"column:password" binding:"required"`
}

func (user *User) Save() error {

	if err := dao.Find("name = ?", getParams(user.Name), user); err != nil {
		return err
	}
	if user.ID != 0 {
		return e.NewError("名称重复,请换个用户名注册！")
	}
	return dao.Insert(user)

}

func (user *User) Login() error {

	err := dao.Find("name = ? and password = ?", getParams(user.Name, user.Password), user)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if user.ID == 0 {
		return e.NewError("用户不存在")
	}
	return nil

}
