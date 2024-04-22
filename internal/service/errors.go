package service

import "errors"

var (
	ErrUserAlreadyExists   = errors.New("用户名已存在")
	ErrCreateUserFailed    = errors.New("创建用户失败")
	ErrGenerateTokenFailed = errors.New("生成令牌失败")
	ErrHashPasswordFailed  = errors.New("密码加密失败")
	ErrInvalidCredentails  = errors.New("用户名或密码错误")

	ErrDishAlreadyExists = errors.New("菜品已存在")
)
