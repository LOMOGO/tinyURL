package utils

import (
	"github.com/jordan-wright/email"
	"time"
	"tinyURL/global"
)

func SendEmail(to []string, subject string, message []byte) error {
	from := global.EMAILCONF.From
	e := email.NewEmail()
	e.From = from
	e.To = to
	e.Subject = subject
	e.Text = message
	timeout := time.Duration(global.EMAILCONF.Timeout) * time.Second
	err := global.EMAILPOOL.Send(e, timeout)
	if err != nil {
		return err
	}
	return nil
}
