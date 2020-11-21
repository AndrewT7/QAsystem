package api

import (
	"github.com/SSunSShine/QAsystem/middleware"
	"github.com/SSunSShine/QAsystem/model"
	"github.com/SSunSShine/QAsystem/service"
	"github.com/SSunSShine/QAsystem/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type UserVO struct {
	ID        uint      `json:"id"`
	Mail      string    `json:"mail"`
	Password  string    `json:"password"`
	Profile   ProfileVO `json:"profile"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetUser 获取单个用户信息
func GetUser(c *gin.Context) {

	var u model.User

	id, _ := strconv.Atoi(c.Param("id"))
	u.ID = uint(id)

	user, err := u.Get()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}
	var userVO UserVO
	util.SimpleCopyProperties(&userVO, &user)
	util.SimpleCopyProperties(&userVO.Profile, &user.Profile)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": userVO,
	})

}

// UpdateUser 更新用户信息。 注：更新Profile个人简介需要调用UpdateProfile
func UpdateUser(c *gin.Context)  {

	var u model.User

	id, _ := strconv.Atoi(c.Param("id"))
	u.ID = uint(id)

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
	}


	if code, err := u.Update(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": code,
		})
	}
}

// DeleteUser
func DeleteUser(c *gin.Context)  {
	
	var u model.User

	id, _ := strconv.Atoi(c.Param("id"))
	u.ID = uint(id)

	user, err := u.Get()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	if code, err := user.Delete(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": code,
		})
	}
}

// GetUsersCount
func GetUsersCount(c *gin.Context)  {

	var u model.User

	if count, err := u.Count(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": count,
		})
	}
}

// Sign 用户注册
func Sign(c *gin.Context)  {

	var s service.SignInterface

	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	if code, err := s.Sign(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": code,
		})
	}
}

// Login 用户登录
func Login(c *gin.Context)  {

	var l service.LoginInterface

	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	code, user, err := l.Login()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	token, err := middleware.Gen(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": code,
		"token": token,
	})
}