package utils

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func NewLogger(lvl logrus.Level) (*logrus.Logger, error) {
	var jsonFormater = &logrus.JSONFormatter{CallerPrettyfier: prettyfier}
	var log = logrus.New()
	log.SetFormatter(jsonFormater)
	log.SetReportCaller(true)
	log.SetLevel(lvl)

	return log, nil
}

func prettyfier(f *runtime.Frame) (string, string) {
	_, filename := path.Split(f.File)
	filename = fmt.Sprintf("%s:%d", filename, f.Line)
	return "", filename
}
