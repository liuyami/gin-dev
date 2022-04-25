package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hanya-gin/app/common/request"
	"hanya-gin/app/common/response"
	"hanya-gin/app/services"
)

func Login(c *gin.Context) {
	var form request.Login

	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.Fail(c, 1001, err.Error())
	} else {
		tokeData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)

		if err != nil {
			response.Fail(c, 1002, err.Error())
		}

		response.Success(c, tokeData)
	}

}

func Register(c *gin.Context) {
	var form request.Register

	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))

	if err != nil {
		response.Fail(c, 4001, "登出失败")
	}

	response.Success(c, nil)
}

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))

	if err != nil {
		response.Fail(c, 4301, err.Error())
	}

	response.Success(c, user)
}
