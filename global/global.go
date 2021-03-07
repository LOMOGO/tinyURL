package global

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tinyURL/config"
)

var (
	ROUTER       *gin.Engine
	DataBaseConf config.Database
	AppConf      config.App
	ZapConf      config.Zap
	QRCodeConf   config.QRCode
	DB           *gorm.DB
	Logger       *zap.Logger
)
