package global

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tinyURL/model"
)

var (
	ROUTER       *gin.Engine
	DataBaseConf model.Database
	AppConf      model.App
	DB           *gorm.DB
)
