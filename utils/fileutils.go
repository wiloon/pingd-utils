package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

func IsFileOrDirExists(path string) bool {

	_, err := os.Stat(path)
	fileExist := err == nil || os.IsExist(err)

	logrus.Infof("file: %s, exist:%v", path, fileExist)

	return false
}
