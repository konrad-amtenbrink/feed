package logger

import (
	"os"

	"github.com/konrad-amtenbrink/feed/internal/config"
	log "github.com/sirupsen/logrus"
)

func MustInit(cfg config.LoggingConfig) {
	level, err := log.ParseLevel(cfg.Level)
	if err != nil {
		log.WithError(err).Fatal("failed to parse logger level")
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(level)
	log.SetReportCaller(cfg.EnableReportCaller)

	if cfg.EnableTextFormatter {
		log.SetFormatter(&log.TextFormatter{})
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}
}
