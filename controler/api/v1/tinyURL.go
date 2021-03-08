package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tinyURL/global"
	"tinyURL/model"
	"tinyURL/model/dao"
	"tinyURL/model/request"
	"tinyURL/utils"
	"tinyURL/utils/errCode"
	"tinyURL/utils/qrcode"
)

// @Summary 生成短链
// @Description 将用户传入的网址转换为短链接, 链接必须以：http://或https://开头
// @Tags 短链
// @Accept json
// @Produce json
// @Param data body request.URLInfo true "需要转换的网址"
// @Success 200 {string} json "{"code": "0", "msg": "成功", "details": "", data:""}"
// @Router /link [post]
func Long2ShortURL(c *gin.Context) {
	var u request.URLInfo
	err := c.BindJSON(&u)
	ok, msg := utils.Valid(err)
	if !ok {
		errCode.InvalidParams.WithDetails(msg).ResponseJson(c)
		global.LOGGER.Warn("入参错误", zap.Error(err))
		return
	}

	url := model.URLInfo{
		Model:   gorm.Model{},
		URL:     u.URL,
		URLCode: "",
	}
	var urlCode string
	urlCode, err = dao.GenderURLCode(&url)
	if err != nil {
		errCode.GenderURLError.ResponseJson(c)
		global.LOGGER.Error("短链生成失败", zap.Error(err))
		return
	}
	tinyUrl := fmt.Sprintf("http://%s:%d/l/%s", global.APPCONF.Host, global.APPCONF.Port, urlCode)
	errCode.Success.WithData(tinyUrl).ResponseJson(c)
	global.LOGGER.Info("短链生成成功", zap.String("短链", urlCode))
}

// @Summary 生成二维码
// @Description 将用户传入的网址转换为二维码，链接必须以：http://或https://开头
// @Tags 二维码
// @Accept json
// @Produce json
// @Param data body request.URLInfo true "需要转换的网址"
// @Success 200 {string} json "{"code": "0", "msg": "成功", "details": "", data:""}"
// @Router /qrcode [post]
func MakeQRCode(c *gin.Context) {
	var u request.URLInfo
	err := c.BindJSON(&u)
	ok, msg := utils.Valid(err)
	if !ok {
		errCode.InvalidParams.WithDetails(msg).ResponseJson(c)
		global.LOGGER.Warn("入参错误", zap.Error(err))
		return
	}

	//生成二维码
	var name string
	name, err = qrcode.GenerateQrcode(u.URL)
	if err != nil {
		errCode.GenderQRCodeError.ResponseJson(c)
		global.LOGGER.Error("二维码生成失败", zap.Error(err))
		return
	}
	qRCodeURL := fmt.Sprintf("http://%s:%d/api/v1/qrcode/%s", global.APPCONF.Host, global.APPCONF.Port, name)
	errCode.Success.WithData(qRCodeURL).ResponseJson(c)
	global.LOGGER.Info("二维码生成成功", zap.String("图片名", name))
}

// @Summary 发送邮件
// @Description 将生成的二维码和短链接发送到用户邮箱，链接必须以：http://或https://开头。
// @Tags 邮件
// @Accept json
// @Produce json
// @Param data body request.EmailInfo true "需要转换的网址"
// @Success 200 {string} json "{"code": "0", "msg": "成功", "details": "", data:""}"
// @Router /email [post]
func SendEmail(c *gin.Context) {
	var e request.EmailInfo
	err := c.BindJSON(&e)
	ok, msg := utils.Valid(err)
	if !ok {
		errCode.InvalidParams.WithDetails(msg).ResponseJson(c)
		global.LOGGER.Error("入参错误", zap.Error(err))
		return
	}
	//生成短链接
	url := model.URLInfo{
		Model:   gorm.Model{},
		URL:     e.URL,
		URLCode: "",
	}
	var urlCode string
	urlCode, err = dao.GenderURLCode(&url)
	if err != nil {
		errCode.GenderURLError.ResponseJson(c)
		global.LOGGER.Error("短链生成失败", zap.Error(err))
		return
	}
	tinyUrl := fmt.Sprintf("http://%s:%d/l/%s", global.APPCONF.Host, global.APPCONF.Port, urlCode)
	global.LOGGER.Info("短链生成成功", zap.String("短网址", tinyUrl))

	//生成二维码
	var name string
	name, err = qrcode.GenerateQrcode(e.URL)
	if err != nil {
		errCode.GenderQRCodeError.ResponseJson(c)
		global.LOGGER.Error("二维码生成失败", zap.Error(err))
		return
	}
	qRCodeURL := fmt.Sprintf("http://%s:%d/api/v1/qrcode/%s", global.APPCONF.Host, global.APPCONF.Port, name)
	global.LOGGER.Info("二维码生成成功", zap.String("图片名", name))

	message := []byte(
		fmt.Sprintf("网址：%s\n"+
			"对应的二维码下载地址为（跳转后长按图片保存或右键鼠标保存）：%s\n"+
			"对应的短链接是：%s", e.URL, qRCodeURL, tinyUrl),
	)

	//发件人和邮件标题
	to := []string{e.Recipient}
	subject := "长网址对应二维码及短链接地址"
	//发送邮件
	err = utils.SendEmail(to, subject, message)
	if err != nil {
		errCode.SendEmailError.WithDetails(err).ResponseJson(c)
		global.LOGGER.Error("邮件发送失败", zap.Error(err), zap.Strings("收件人:", to), zap.Binary("邮件信息为：", message))
		return
	}
	errCode.Success.WithData(gin.H{"收件人": to}).ResponseJson(c)
	global.LOGGER.Info("邮件发送成功", zap.Strings("收件人:", to), zap.Binary("邮件信息为：", message))
}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")
	var url model.URLInfo
	url.URLCode = code
	err := dao.SelectByCode(&url)
	if err != nil {
		errCode.NotFound.ResponseJson(c)
		global.LOGGER.Warn("未找到短链对应的网址", zap.Error(err))
		return
	}
	c.Redirect(302, url.URL)
	global.LOGGER.Info("网页跳转成功", zap.String("短链", code), zap.String("跳转的网址：", url.URL))
}
