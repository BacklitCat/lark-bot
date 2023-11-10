package config

var Bot BotConfig

type BotConfig struct {
	Name   string `yaml:"Name"`
	Server Server `yaml:"Server"`
	Lark   Lark   `yaml:"Lark"`
}

type Server struct {
	Port string `yaml:"Port"`
}

type Lark struct {
	AppID                      string `yaml:"AppID"`
	Secret                     LarkSecret
	UpdateTokenDeltaTimeSecond int64 `yaml:"UpdateTokenDeltaTimeSecond"`
}

func init() {
	MustLoadConfig(ProjectPath+"/etc/lark_secret.yml", &LSecret)
	MustLoadConfig(ProjectPath+"/etc/bot_config.yml", &Bot)
	Bot.Lark.Secret = LSecret
}
