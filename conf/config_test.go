package conf

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Println(GetConfig().RoomConfig)
}