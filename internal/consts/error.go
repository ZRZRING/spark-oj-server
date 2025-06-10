package consts

import "github.com/gogf/gf/v2/errors/gerror"

// Error Message
var (
	ErrUserExist       = gerror.New("用户已存在")
	ErrUserNotExist    = gerror.New("用户不存在")
	ErrInvalidPassword = gerror.New("密码错误")
	ErrInvalidPageSize = gerror.New("错误的分页数据")
	ErrProblemNotExist = gerror.New("题目不存在")
)
