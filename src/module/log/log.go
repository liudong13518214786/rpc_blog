package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

//todo 日志(m目前暂时使用beego日志模块)
var log = logrus.New()

type Log struct {
	serviceName string
}

func (l *Log) Init() *logrus.Logger {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	log.WithFields(logrus.Fields{
		"serverName": l.serviceName,
	})
	log.Infof("init %v service log...", l.serviceName)
	return log
}

func (l *Log) Info(message string) {
	log.Info(message)
}

func (l *Log) Error(message string) {
	log.Error(message)
}
