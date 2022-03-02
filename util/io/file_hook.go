package io

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

///文件日志Hook
type FileLoggerHook struct {
	maxLogSize int64
	logDir     string
	logPath    string
	fileHeader   string
	logFile    *os.File
}

///工厂方法
func NewFileLoggerHook(logDir string, maxLogSize int64,fileHeader string) (hook *FileLoggerHook, err error) {
	object := &FileLoggerHook{
		maxLogSize: maxLogSize,
		logDir:     logDir,
		fileHeader: fileHeader,
	}
	return object, object.makeFile()
}

///创建文件
func (object *FileLoggerHook) makeFile() error {
	var err error
	if filepath.IsAbs(object.logDir) {
		object.logPath = object.logDir
	} else {
		dir, err := os.Getwd()
		if nil != err {
			panic(err)
		}
		object.logPath = filepath.Join(dir, object.logDir)
	}
	err = os.MkdirAll(object.logDir, 0777)
	if nil != err {
		panic(err)
	}

	now := time.Now()
	object.logPath += fmt.Sprintf("%s%s-%04d-%02d-%02d.dat",
		string(filepath.Separator),object.fileHeader, now.Year(), now.Month(), now.Day())
	if nil != object.logFile {
		err = object.logFile.Close()
		if nil != err {
			return err
		}
	}
	object.logFile, err = os.OpenFile(object.logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if nil != err {
		return err
	}
	return nil
}

func (object *FileLoggerHook) String() string {
	return object.logPath
}

