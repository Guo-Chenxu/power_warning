package logic

import (
	"power_warning/conf"
	"testing"
)

func TestGetPower(t *testing.T) {
	p, err := GetPower(conf.GetConfig().RoomConfig)
	if err != nil {
		t.Error(err)
	}
	t.Log(p.D.Data.Surplus)
}
