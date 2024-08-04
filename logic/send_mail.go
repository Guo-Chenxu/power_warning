package logic

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"power_warning/conf"

	"github.com/jordan-wright/email"
)

func SendEmail(mailConfig conf.MailConfig) error {
	auth := smtp.PlainAuth("", mailConfig.From, mailConfig.Secret, mailConfig.Host)
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s<%s>", mailConfig.Nickname, mailConfig.From)
	e.To = mailConfig.To
	e.Subject = mailConfig.Subject
	e.HTML = []byte(mailConfig.Body)
	hostAddr := fmt.Sprintf("%s:%d", mailConfig.Host, mailConfig.Port)
	if mailConfig.Ssl {
		return e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: mailConfig.Host})
	}
	return e.Send(hostAddr, auth)
}
