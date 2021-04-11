package bizerr

import "errors"

var (
	ErrDB = errors.New("操作数据出错")

	ErrUserNotFound   = errors.New("用户未找到")
	ErrUserInvalidPwd = errors.New("密码无效")
)
