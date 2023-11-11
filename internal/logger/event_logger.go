package logger

import (
	"larkbot/internal/config"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var EventLogger *log.Logger

func init() {

	if err := os.MkdirAll(config.LogPath, 0755); err != nil {
		log.Fatalln("fatal: make log path faild")
	}

	EventLogger = log.New()

	fileName := path.Join(config.LogPath, "larkbot.event")

	EventLogger.AddHook(NewLfsHookDayLevel(fileName))
}

func OpenLogFile(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func NewRotateDayWriter(fileName string) *rotatelogs.RotateLogs {
	rotateWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logrus.Fatalf("Failed to create rotateWriter: %v", err)
	}
	return rotateWriter
}

func NewLfsHookDayLevel(fileName string) *lfshook.LfsHook {
	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: NewRotateDayWriter(fileName + ".debug"),
		logrus.InfoLevel:  NewRotateDayWriter(fileName + ".info"),
		logrus.WarnLevel:  NewRotateDayWriter(fileName + ".warn"),
		logrus.ErrorLevel: NewRotateDayWriter(fileName + ".error"),
		logrus.FatalLevel: NewRotateDayWriter(fileName + ".fatel"),
		logrus.PanicLevel: NewRotateDayWriter(fileName + ".panic"),
	}
	return lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
