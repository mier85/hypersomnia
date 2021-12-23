package main

import (
	"github.com/kstkn/hypersomnia/server"
	log "github.com/sirupsen/logrus"

	"github.com/kstkn/hypersomnia/config"
)

type LogrusWrapper struct {
	*log.Logger
}

func (l *LogrusWrapper) Warn(message string, fields ...[2]string) {
	logFields := make(log.Fields, len(fields))
	for _, msg := range fields {
		logFields[msg[0]] = msg[1]
	}
	l.WithFields(logFields).Warn(message)
}

func main() {
	conf := config.NewConfig()
	log.SetLevel(conf.LogLevel)
	log.Debugf("configuration %+v", conf)
	server.StartServer(&LogrusWrapper{log.StandardLogger()}, conf)
}
