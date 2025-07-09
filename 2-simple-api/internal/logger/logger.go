package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	Log.SetLevel(logrus.DebugLevel)
	Log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05:000"})
	Log.SetOutput(os.Stdout)
}
