package logger

import (
	"os"
	"time"

	"go-boilerplate-app/pkg/utils/constants"

	level "github.com/labstack/gommon/log"
	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
)

type LogLevel string

var appLogger *log.MyLogger

func init() {
	log.Logger().SetOutput(os.Stdout)
	// Default log level
	log.Logger().SetLevel(getLogLevel(constants.DEFAULT_LOG_LEVEL))
	log.Logger().SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		DisableColors:   false,
	})
	// log.Logger().SetReportCaller(true)
	appLogger = log.Logger()
}

func SetLogger(lvl string) {
	appLogger.SetLevel(getLogLevel(lvl))
}

func SetDevMode() {
	SetLogger(constants.DEFAULT_DEV_LOG_LEVEL)
}

func GetLogger() *log.MyLogger {
	return appLogger
}

func Debug(msg string, args ...interface{}) {
	appLogger.Logf(logrus.DebugLevel, msg, args...)
}

func Info(msg string, args ...interface{}) {
	appLogger.Logf(logrus.InfoLevel, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	appLogger.Logf(logrus.WarnLevel, msg, args...)
}

func Error(msg string, args ...interface{}) {
	appLogger.Logf(logrus.ErrorLevel, msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	appLogger.Logf(logrus.FatalLevel, msg, args...)
}

func Panic(msg string, args ...interface{}) {
	appLogger.Logf(logrus.PanicLevel, msg, args...)
}

const (
	DebugLvl LogLevel = "debug"
	InfoLvl  LogLevel = "info"
	WarnLvl  LogLevel = "warn"
	ErrorLvl LogLevel = "error"
	FatalLvl LogLevel = "fatal"
	PanicLvl LogLevel = "panic"
)

func getLogLevel(lvl string) level.Lvl {
	switch lvl {
	case string(DebugLvl):
		return 1
	case string(InfoLvl):
		return 2
	case string(WarnLvl):
		return 3
	case string(ErrorLvl):
		return 4
	case string(FatalLvl):
		return 7
	case string(PanicLvl):
		return 6
	default:
		return 5
	}
}
