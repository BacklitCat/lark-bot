package config_test

import (
	"larkbot/internal/config"
	"os"
	"testing"
)

func TestLoadBotConfig(t *testing.T) {
	projectPath := os.Getenv("PROJECT_LARK_BOT")
	configFilePath := projectPath + "/etc/bot_config.yml"
	t.Logf("configFilePath=%s\n", configFilePath)

	var conf config.BotConfig

	config.MustLoadConfig(configFilePath, &conf)

	t.Log(conf)
}
