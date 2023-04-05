/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"time"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

type Logger struct{ *logrus.Logger }

func (l *Logger) Say(msg string) {
	l.Info(msg)
}

func (l *Logger) Sayf(fmt string, args ...interface{}) {
	l.Infof(fmt, args...)
}

func (l *Logger) SayWithField(msg string, k string, v map[string]string) {
	l.WithField(k, v).Info(msg)
}
func (l *Logger) SayWithFields(msg string, fields map[string]interface{}) {
	l.WithFields(fields).Info(msg)
}

func StdOutFileLogger(outputFileName, timeStampFormat string, maxSize, maxBackups, maxAge int) *Logger {

	logLevel := logrus.InfoLevel
	log := logrus.New()
	log.SetLevel(logLevel)
	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   outputFileName,
		MaxSize:    maxSize,    // megabytes
		MaxBackups: maxBackups, // amouts
		MaxAge:     maxAge,     //days
		Level:      logLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC822,
		},
	})

	if err != nil {
		logrus.Fatalf("Failed to initialize file rotate hook: %v", err)
	}
	log.SetOutput(colorable.NewColorableStdout())

	log.SetFormatter(&logrus.TextFormatter{
		PadLevelText:    true,
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: timeStampFormat,
	})

	log.AddHook(rotateFileHook)

	return &Logger{log}
}
