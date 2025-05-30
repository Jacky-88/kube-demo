package service

import (
	"errors"
	"github.com/wonderivan/logger"
	"k8s-demo-test/config"
)

var Login login

type login struct {
}

// 登录验证
func (l *login) Auth(username, password string) (err error) {
	if username == config.AdminUser && password == config.AdminPwd {
		return nil
	} else {
		logger.Error("用户名或密码错误")
		return errors.New("用户名或密码错误")
	}
}
