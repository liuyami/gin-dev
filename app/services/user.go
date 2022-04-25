package services

import (
	"errors"
	"hanya-gin/app/common/request"
	"hanya-gin/app/models"
	"hanya-gin/global"
	"hanya-gin/utils"
	"strconv"
)

type userService struct {
}

var UserService = new(userService)

func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error

	if err != nil {
		err = errors.New("账户数据不存在")
	}

	return
}

func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.App.DB.Where("mobile=?", params.Mobile).First(&user).Error

	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("错误的账号或密码")
	}

	return

}

func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("mobile=?", params.Mobile).Select("id").First(&models.User{})

	if result.RowsAffected != 0 {
		err = errors.New("手机号码已注册过")
		return
	}

	user = models.User{
		Name:     params.Name,
		Mobile:   params.Mobile,
		Openid:   params.Openid,
		Password: utils.BcryptMake([]byte(params.Password)),
	}

	err = global.App.DB.Create(&user).Error

	return
}
