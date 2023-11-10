package config_test

import (
	"larkbot/internal/config"
	"os"
	"testing"
)

func TestLoadAPIConfig(t *testing.T) {
	projectPath := os.Getenv("PROJECT_LARK_BOT")
	configFilePath := projectPath + "/etc/larkapi_config.yml"
	t.Logf("configFilePath=%s\n", configFilePath)

	var conf config.LarkAPIConfig

	config.MustLoadConfig(configFilePath, &conf)

	if conf.APP_ACCESS_TOKEN.APIName == "" {
		t.Error("TestLoadAPIConfig faild: can't get conf")
	}

	t.Log(conf)
}
