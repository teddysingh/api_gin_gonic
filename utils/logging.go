package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger - common logger to use across application
var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.Level = logrus.DebugLevel
	Logger.Out = os.Stdout
}
