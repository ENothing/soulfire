package logging

import (
	"soulfire/utils"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

var (
	FileSavePath = "runtime/logs"
	LevelFlags    = []string{"DEBUG", "INFO", "WARN", "ERR", "FATAL"}
	logger *log.Logger
)

const (
	DEBUG int = iota
	INFO
	WARN
	ERR
	FATAL
)

func init() {

	filename := time.Now().Format("20060102") + ".log"

	logfile := utils.OpenFile(FileSavePath,filename)

	logger = log.New(logfile,"",log.LstdFlags)
}

func Logging(logtype int,v ...interface{}){

	var prefix string
	_,file,line,ok := runtime.Caller(2)

	if ok {
		prefix = "["+LevelFlags[logtype]+"]" + "["+filepath.Base(file) +":"+ strconv.Itoa(line) +"]"
	}else{
		prefix = "["+LevelFlags[logtype]+"]"
	}

	logger.SetPrefix(prefix)
	logger.Println(v)

}
