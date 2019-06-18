package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

//todo 日志
var log = logrus.New()

func Init() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	log.WithFields(logrus.Fields{
		"serverName": "1231",
	})
	log.Debug("123")
}
