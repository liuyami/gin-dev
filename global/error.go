package global

type CustomError struct {
	ErrCode int
	ErrMsg  string
}

type CustomErrors struct {
	ValidateError CustomError
	BusinessError CustomError
	TokenError    CustomError
}

var Errors = CustomErrors{
	ValidateError: CustomError{
		ErrCode: 4000,
		ErrMsg:  "请求参数错误",
	},

	BusinessError: CustomError{
		ErrCode: 4100,
		ErrMsg:  "业务错误",
	},

	TokenError: CustomError{
		ErrCode: 4300,
		ErrMsg:  "登录授权失效",
	},
}
