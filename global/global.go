package global

import (
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tinyURL/config"
)

var (
	ROUTER       *gin.Engine
	DATABASECONF config.Database
	APPCONF      config.App
	ZAPCONF      config.Zap
	QRCODECONF   config.QRCode
	EMAILCONF    config.Email
	DB           *gorm.DB
	LOGGER       *zap.Logger
	EMAILPOOL    *email.Pool
)
