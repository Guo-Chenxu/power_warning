package logic

import (
	"power_warning/conf"
	"testing"
)

func TestSendMail(t *testing.T) {
	err := SendEmail(conf.GetConfig().MailConfig)
	if err != nil {
		t.Error(err)
	}
}
