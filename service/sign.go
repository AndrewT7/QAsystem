package service

import (
	"errors"
	"github.com/SSunSShine/QAsystem/database"
	"github.com/SSunSShine/QAsystem/model"
	"github.com/SSunSShine/QAsystem/util"
	"log"
)

var validchan = make(chan string, 10)

// SignInterface 注册接口
type SignInterface struct {
	Name     string `json:"name" validate:"required,min=4,max=16" label:"用户名"`
	Gender   int    `json:"gender"`
	Mail     string `json:"mail" validate:"required,email" label:"邮箱号"`
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
	Phone    string `json:"phone" validate:"required,min=11,max=11" label:"手机号"`
}

//验证码校验接口
type SignInterfaceCode struct {
	Code  string `json:"code" validate:"required,min=6,max=6" label:"验证码"`
	Phone string `json:"phone" validate:"required,min=11,max=11" label:"手机号"`
}

func (s *SignInterface) Sign() (err error) {

	var p model.Profile
	var u model.User

	msg, err := util.Validate(s)
	if err != nil {
		log.Println(msg)
		return errors.New(msg)
	}

	u.Mail = s.Mail
	u.Password = s.Password
	u.Phone = s.Phone
	util.SendCode(u.Phone)

	//检验验证码是否正确
	val := <-validchan

	if val == s.Phone {
		//需要判断验证码
		err = u.Create()
		if err != nil {
			return
		}

		p.Name = s.Name
		p.Desc = "Hello ~ " + p.Name
		p.Gender = s.Gender
		p.UserID = u.ID
		err = p.Create()
		if err != nil {
			return
		}
		return nil
	}
	return errors.New("用户注册失败")
}

func (s *SignInterfaceCode) Validcode() (err error) {
	msg, err := util.Validate(s)
	if err != nil {
		log.Println(msg)
		return errors.New(msg)
	}
	code, _ := database.RDB.Get(s.Phone).Result()
	if code == s.Code {
		validchan <- s.Phone
	}

	return
}
