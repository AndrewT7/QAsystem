package service

import (
	"errors"
	"github.com/SSunSShine/QAsystem/model"
	"github.com/SSunSShine/QAsystem/util"
	"golang.org/x/crypto/bcrypt"
	"log"

)

// LoginInterface 登录接口
type LoginInterface struct {
	Phone	 string `json:"phone" validate:"required,min=11,max=11" label:"手机号"`
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
}

func (l *LoginInterface) Login() (user model.User, err error) {

	var u model.User

	msg, err := util.Validate(l)
	if err != nil {
		log.Println(msg)
		return u, errors.New(msg)
	}

	u.Phone = l.Phone
	user, err = u.Get()
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(l.Password)); err != nil {
		log.Print(err)
		return
	}

	return
}

