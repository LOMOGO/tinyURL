package initialize

import (
	"github.com/gin-gonic/gin"
	"tinyURL/global"
)

func Routers() {
	global.ROUTER = gin.Default()
}
