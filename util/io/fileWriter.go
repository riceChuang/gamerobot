package io

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const (
	defaultFilePath = "outputs"
	maxFileSize    = 10 << 20 //最大日志文件大小10MB
	maxDirSize     = 1 << 30  //最大日志文件目录大小1GB
)

var (
	fileWriter *FileWriter
)


//建立file方式 -> outputs/"FolderName"/FileHeader-yyyy-mm-dd.dat

type FileConf struct {
	FolderName string
	FileHeader string    //文件開頭
}

type FileWriter struct {
	cfg       *FileConf
	file      *os.File
	hook      *FileLoggerHook
	writeChan chan string
	stopChan  chan bool
}

func NewFileWriter(cfg *FileConf) (*FileWriter, error) {
	if cfg.FolderName == "" {
		return nil, errors.New("FolderName can't be empty")
	}

	if cfg.FileHeader == "" {
		return nil, errors.New("FileHeader can't be empty")
	}

	dir := fmt.Sprintf("%v/%v/",defaultFilePath,cfg.FolderName)
	fileHook,err := NewFileLoggerHook(dir,maxFileSize,cfg.FileHeader)

	writer := &FileWriter{
		cfg:       cfg,
		file:      fileHook.logFile,
		writeChan: make(chan string, 10000),
		stopChan:  make(chan bool, 1),
	}
	writer.run()
	return writer, err
}

func (f *FileWriter) run() {
	go f.writeFile()
}

func (f *FileWriter) Write(data string) {
	f.writeChan <- data
}

func (f *FileWriter) Stop() {
	f.stopChan <- true

	//清空channel
	for len(f.writeChan) > 0 {
		data := <-f.writeChan
		_, err := fmt.Fprintln(f.file, data)
		if err != nil {
			fmt.Printf("write file fail for file %v,error = %v\n", f.hook.String(), err)
		}
	}
	defer f.file.Close()
	err := f.file.Close()
	if err != nil {
		fmt.Printf("fail to stop file writer err= %v\n", err)
	}
}

func (f *FileWriter) writeFile() {
outer:
	for {
		select {
		case data := <-f.writeChan:
			{
				_, err := f.file.WriteString(data + "\n")
				if err != nil {
					fmt.Printf("write file fail for file %v,error = %v\n", f.hook.String(), err)
				}
			}
		case <-f.stopChan:
			{
				//關閉write channel
				close(f.writeChan)
				break outer
			}
		default:
			{
				time.Sleep(time.Second * 1)
			}
		}
	}
}
