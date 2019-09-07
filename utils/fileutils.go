package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

func IsDirExists(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		logrus.Error(e)
		return false
	}
	if f != nil && f.IsDir() {
		return true
	}

	return false
}
