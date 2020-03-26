package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var out = os.Stdout
var logger = logrus.New()

func init() {
	logger.SetOutput(out)
	logger.SetFormatter(&textFormatter{})

	httpLogger = logrus.New()
	httpLogger.SetOutput(out)
	httpLogger.Level = logrus.DebugLevel
	httpLogger.SetFormatter(&textFormatter{})
}

func Info(logs ...interface{}) {
	logger.Info(logs...)
}

func Fatal(logs ...interface{}) {
	logger.Fatal(logs...)
}

func Warn(logs ...interface{}) {
	logger.Warn(logs...)
}

func Error(logs ...interface{}) {
	logger.Error(logs...)
}
