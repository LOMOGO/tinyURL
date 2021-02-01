package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"tinyURL/global"
	"tinyURL/model"
	"tinyURL/model/dao"
	"tinyURL/utils/errCode"
)

func Long2ShortURL(c *gin.Context) {
	var url *model.URLInfo
	err := c.BindJSON(&url)
	if err != nil {
		log.Println("post参数绑定错误:", err)
		c.JSON(errCode.BindError.StatusCode(), errCode.BindError)
		return
	}
	urlCode, err := dao.GenderURLCode(url)
	if err != nil {
		c.JSON(errCode.GenderURLError.StatusCode(), errCode.GenderURLError)
		return
	}
	success := errCode.Success
	success.Data = "http://" + global.AppConf.URL + global.AppConf.Port + "/" + urlCode
	c.JSON(errCode.Success.StatusCode(), success)
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
