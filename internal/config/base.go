package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ProjectPath = os.Getenv("PROJECT_LARK_BOT")
	LogPath     = ProjectPath + "/log"
)

func MustLoadConfig(filePath string, conf any) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("[FATAl] func=MustLoadConfig, step=os.ReadFile, filePath=%s, err:%s\n", filePath, err.Error())
	}

	if len(content) == 0 {
		log.Fatalf("[FATAl] func=MustLoadConfig, step=os.ReadFile, filePath=%s, err:%s\n", filePath, "file content empty")
	}
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		log.Fatalf("[FATAl] func=MustLoadConfig, step=yaml.Unmarshal, filePath=%s, content=%v, err:%s\n", filePath, conf, err.Error())
	}
}
