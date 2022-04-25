package controllers

import (
	"github.com/gin-gonic/gin"
	"hanya-gin/app/common/request"
	"hanya-gin/app/common/response"
	"hanya-gin/app/services"
)

func UploadImage(c *gin.Context) {
	var form request.ImageUpload
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	outPut, err := services.MediaService.SaveImage(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, outPut)
}
