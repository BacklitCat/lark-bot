package config

var LarkAPI LarkAPIConfig

type DefaultLarkAPI struct {
	APIName string `yaml:"name"`
	Url     string `yaml:"url"`
	Method  string `yaml:"method"`
}

type LarkAPIConfig struct {
	APP_ACCESS_TOKEN    DefaultLarkAPI `yaml:"app_access_token"`
	TENANT_ACCESS_TOKEN DefaultLarkAPI `yaml:"tenant_access_token"`
	MESSAGES            DefaultLarkAPI `yaml:"messages"`
}

func init() {
	MustLoadConfig(ProjectPath+"/etc/larkapi_config.yml", &LarkAPI)
}
