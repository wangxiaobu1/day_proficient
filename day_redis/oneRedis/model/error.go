package model

import (
	"errors"
)

//根据业务逻辑需要，自动以一些错误。
var (
	ERROR_USER_NOTEXITS = errors.New("用户不存在")
	ERROR_USER_EXITS = errors.New("用户已存在")
	ERROR_USER_PWDNOTTRUE = errors.New("用户名或者密码不正确")
)
