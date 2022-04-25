package response

import (
	"github.com/gin-gonic/gin"
	"hanya-gin/global"
	"net/http"
	"os"
)

type Response struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		ErrCode: 0,
		ErrMsg:  "ok",
		Data:    data,
	})
}

func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response{
		ErrCode: errorCode,
		ErrMsg:  msg,
		Data:    nil,
	})
}

func FailByError(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrCode, error.ErrMsg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidateError.ErrCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.BusinessError.ErrCode, msg)
}

func TokenFail(c *gin.Context) {
	FailByError(c, global.Errors.TokenError)
}

func ServerError(c *gin.Context, err any) {
	msg := "Internal Server Error"

	// 非生产环境显示具体错误信息
	if global.App.Config.App.Env != "production" && os.Getenv(gin.EnvGinMode) != gin.ReleaseMode {
		if _, ok := err.(error); ok {
			msg = err.(error).Error()
		}
	}

	c.JSON(http.StatusInternalServerError, Response{
		http.StatusInternalServerError,
		msg,
		nil,
	})

	c.Abort()
}
