package middleware

import (
	"larkbot/internal/config"
	"larkbot/internal/logger"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerToFile() gin.HandlerFunc {

	midlogger := logrus.New()

	fileName := path.Join(config.LogPath, "larkbot.gin")

	midlogger.AddHook(logger.NewLfsHookDayLevel(fileName))

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		midlogger.WithFields(logrus.Fields{
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
