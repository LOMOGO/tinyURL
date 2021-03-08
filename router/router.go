package router

import (
	"fmt"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	v1 "tinyURL/controler/api/v1"
	"tinyURL/global"
	"tinyURL/utils/qrcode"
)

func RunAppServer() {
	global.ROUTER.GET("/l/:code", v1.RedirectURL)
	global.ROUTER.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	version1 := global.ROUTER.Group("/api/v1")
	version1.POST("/link", v1.Long2ShortURL)
	version1.POST("/qrcode", v1.MakeQRCode)
	version1.POST("/email", v1.SendEmail)
	//这里只要把二维码的存放路径写进去即可，访问时传入图片名称即可获取图片
	version1.StaticFS("/qrcode", http.Dir(qrcode.GetQRCodeFullPath()))

	port := fmt.Sprintf(":%d", global.APPCONF.Port)
	global.ROUTER.Run(port)
}
