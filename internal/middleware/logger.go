package middleware

import (
	"fmt"
	"larkbot/internal/config"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func LoggerToFile() gin.HandlerFunc {

	logger := logrus.New()

	fileName := path.Join(config.LogPath, "larkbot.gin")

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	logger.Out = src

	logger.SetLevel(logrus.InfoLevel)

	rotateWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	if err != nil {
		logrus.Fatalf("Failed to create logWriter: %v", err)
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  rotateWriter,
		logrus.FatalLevel: rotateWriter,
		logrus.DebugLevel: rotateWriter,
		logrus.WarnLevel:  rotateWriter,
		logrus.ErrorLevel: rotateWriter,
		logrus.PanicLevel: rotateWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}

// 日志记录到 ES，待写
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
