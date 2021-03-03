package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tinyURL/global"
	"tinyURL/model"
	"tinyURL/model/dao"
	"tinyURL/model/request"
	"tinyURL/utils"
	"tinyURL/utils/errCode"
)

// @Summary 生成短链
// @Description 将用户传入的网址转换为短链接
// @Tags 短链
// @Accept json
// @Produce json
// @Param data body request.URLInfo true "需要转换的网址"
// @Success 200 {string} json "{"code": "0", "msg": "成功", "details": "", data:"转换后的短链"}"
// @Router /link [post]
func Long2ShortURL(c *gin.Context) {
	var u request.URLInfo
	var msg map[string]string
	err := c.BindJSON(&u)
	ok, msg := utils.Valid(err)
	if !ok {
		errCode.InvalidParams.WithDetails(msg).ResponseJson(c)
		global.Logger.Warn("入参错误", zap.Error(err))
		return
	}

	url := model.URLInfo{
		Model:   gorm.Model{},
		URL:     u.URL,
		URLCode: "",
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
