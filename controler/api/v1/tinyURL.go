package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"tinyURL/global"
	"tinyURL/model"
	"tinyURL/model/dao"
	"tinyURL/utils"
	"tinyURL/utils/errCode"
)

func Long2ShortURL(c *gin.Context) {
	var url model.URLInfo
	var msg map[string]string
	err := c.BindJSON(&url)
	ok, msg := utils.Valid(err)
	if !ok {
		errCode.InvalidParams.WithDetails(msg).ResponseJson(c)
		global.Logger.Warn("入参错误", zap.Error(err))
		return
	}

	urlCode, err := dao.GenderURLCode(&url)
	if err != nil {
		errCode.GenderURLError.ResponseJson(c)
		global.Logger.Error("短链生成失败", zap.Error(err))
		return
	}
	errCode.Success.WithData("http://" + global.AppConf.URL + global.AppConf.Port + "/" + urlCode).ResponseJson(c)
	global.Logger.Info("短链生成成功", zap.String("短链：", urlCode))
}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")
	var url model.URLInfo
	url.URLCode = code
	err := dao.SelectByCode(&url)
	if err != nil {
		errCode.NotFound.ResponseJson(c)
		global.Logger.Warn("未找到短链对应的网址", zap.Error(err))
		return
	}
	c.Redirect(302, url.URL)
	global.Logger.Info("网页跳转成功", zap.String("短链", code), zap.String("跳转的网址：", url.URL))
}
