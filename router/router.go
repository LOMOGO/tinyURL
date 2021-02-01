package router

import (
	v1 "tinyURL/controler/api/v1"
	"tinyURL/global"
)

func RunAppServer() {
	global.ROUTER.GET("/l/:code", v1.RedirectURL)

	version1 := global.ROUTER.Group("/v1")
	version1.POST("/link", v1.Long2ShortURL)

	global.ROUTER.Run(global.AppConf.Port)
}
