package router

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	v1 "tinyURL/controler/api/v1"
	"tinyURL/global"
)

func RunAppServer() {
	global.ROUTER.GET("/l/:code", v1.RedirectURL)
	global.ROUTER.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	version1 := global.ROUTER.Group("/api/v1")
	version1.POST("/link", v1.Long2ShortURL)

	global.ROUTER.Run(global.AppConf.Port)
}
