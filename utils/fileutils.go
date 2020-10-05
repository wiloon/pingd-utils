package utils

import (
	log "github.com/wiloon/pingd-log/v2/logconfig/zaplog"
	"os"
)

func IsFileOrDirExists(path string) bool {

	_, err := os.Stat(path)
	fileExist := err == nil || os.IsExist(err)

	log.Infof("file: %s, exist:%v", path, fileExist)

	return fileExist
}
