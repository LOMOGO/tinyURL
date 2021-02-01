package v1

import (
	"github.com/gin-gonic/gin"
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
		c.JSON(errCode.InvalidParams.StatusCode(), errCode.InvalidParams.WithDetails(msg))
		return
	}

	urlCode, err := dao.GenderURLCode(&url)
	if err != nil {
		c.JSON(errCode.GenderURLError.StatusCode(), errCode.GenderURLError)
		return
	}
	c.JSON(errCode.Success.StatusCode(), errCode.Success.WithData("http://"+global.AppConf.URL+global.AppConf.Port+"/"+urlCode))
}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")
	var url model.URLInfo
	url.URLCode = code
	err := dao.SelectByCode(&url)
	if err != nil {
		c.JSON(errCode.NotFound.StatusCode(), errCode.NotFound)
		return
	}
	c.Redirect(302, url.URL)
}
