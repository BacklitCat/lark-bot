package config_test

import (
	"larkbot/internal/config"
	"testing"
)

func TestInitConfig(t *testing.T) {
	if config.Bot.Lark.AppID == "" {
		t.Errorf("TestInitConfig faild, config.Bot=%v", config.Bot)
	}

	if config.LarkAPI.APP_ACCESS_TOKEN.APIName == "" {
		t.Errorf("TestInitConfig faild, config.LarkAPI.APP_ACCESS_TOKEN=%v", config.LarkAPI.APP_ACCESS_TOKEN)
	}

}
