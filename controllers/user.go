package controllers

import (
	"fyoukuApi/models"
	"regexp"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

//用户注册功能
// @router /register/save [post]
func (this *UserController) SaveRegister() {
	var (
		mobile   string
		password string
		err      error
	)
	mobile = this.GetString("mobile")
	password = this.GetString("password")

	if mobile == "" {
		this.Data["json"] = ReturnError(4001, "手机号不能为空")
		this.ServeJSON()
	}
	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		this.Data["json"] = ReturnError(4002, "手机号格式不正确")
		this.ServeJSON()
	}
	if password == "" {
		this.Data["json"] = ReturnError(4003, "密码不能为空")
		this.ServeJSON()
	}

	//判断手机号是否已经注册
	status := models.IsUserMobile(mobile)
	if status {
		this.Data["json"] = ReturnError(4005, "此手机号已经注册")
		this.ServeJSON()
	} else {
		err = models.UserSave(mobile, MD5V(password))
		if err == nil {
			this.Data["json"] = ReturnSuccess(0, "注册成功", nil, 0)
			this.ServeJSON()
		} else {
			this.Data["json"] = ReturnError(5000, err)
			this.ServeJSON()
		}
	}
}

//用户登录
// @router /login/do [*]
func (this *UserController) LoginDo() {
	mobile := this.GetString("mobile")
	password := this.GetString("password")

	if mobile == "" {
		this.Data["json"] = ReturnError(4001, "手机号不能为空")
		this.ServeJSON()
	}
	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		this.Data["json"] = ReturnError(4002, "手机号格式不正确")
		this.ServeJSON()
	}
	if password == "" {
		this.Data["json"] = ReturnError(4003, "密码不能为空")
		this.ServeJSON()
	}
	uid, name := models.IsMobileLogin(mobile, MD5V(password))
	if uid != 0 {
		this.Data["json"] = ReturnSuccess(0, "登录成功", map[string]interface{}{"uid": uid, "username": name}, 1)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "手机号或密码不正确")
		this.ServeJSON()
	}
}
