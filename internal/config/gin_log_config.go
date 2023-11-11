package config

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

// 使用中间件后，不再需要

func init() {
	// if err := os.MkdirAll(LogPath, 0755); err != nil {
	// 	log.Fatalln("fatal: make log path faild")
	// }

	// InitGinLogger() // 使用中间件后，不再需要
}

// 配置日志切割
// LogFileCut 日志文件切割
func LogFileCut(fileName string) *rotatelogs.RotateLogs {
	logier, err := rotatelogs.New(
		// 切割后日志文件名称
		fileName+".%Y%m%d%H.log",
		rotatelogs.WithLinkName(fileName),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(30*24*time.Hour), // 文件最大保存时间（30Day）
		rotatelogs.WithRotationTime(time.Hour), // 日志切割时间间隔
		//rotatelogs.WithRotationCount(3),
		// rotatelogs.WithRotationTime(time.Second), // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}
	return logier
}

func InitGinLogger() {
	logFileName := path.Join(LogPath, "gin")

	logFileCut := LogFileCut(logFileName)
	writers := []io.Writer{
		logFileCut,
		os.Stdout}

	gin.DefaultWriter = io.MultiWriter(writers...)

}
