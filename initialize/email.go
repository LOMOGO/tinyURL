package initialize

import (
	"fmt"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"net/smtp"
	"tinyURL/global"
)

//为避免频繁与SMTP服务器建立连接而造成性能问题，创建邮件连接池
func EmailPool() {
	p, err := email.NewPool(
		fmt.Sprintf("%s:%d", global.EMAILCONF.Host, global.EMAILCONF.Port),
		global.EMAILCONF.ConnNum,
		smtp.PlainAuth("", global.EMAILCONF.From, global.EMAILCONF.Secret, global.EMAILCONF.Host),
	)
	if err != nil {
		global.LOGGER.Error("邮件连接池创建失败", zap.Error(err))
	}

	global.EMAILPOOL = p
}
