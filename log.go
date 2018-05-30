package logging

import (
	"fmt"
	"os"
)

var logger = MustGetLogger("example")
var formatFile = MustStringFormatter(
	`%{time:01-02 15:04:05.000} %{level:.4s} %{shortpkg}.%{shortfunc} %{message}`,
)
var formatStd = MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{level:.4s} %{shortpkg}.%{shortfunc} %{color:reset} %{message}`,
)

func Init(logfilename string, logFileLevel Level) {
	logFile, err := os.OpenFile(logfilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	backendFile := NewLogBackend(logFile, "", 0)
	backendStd := NewLogBackend(os.Stdout, "", 0)

	backendFileFormatter := NewBackendFormatter(backendFile, formatFile)
	backendStdFormatter := NewBackendFormatter(backendStd, formatStd)

	backendFileLeveled := AddModuleLevel(backendFileFormatter)
	backendFileLeveled.SetLevel(logFileLevel, "")

	SetBackend(backendFileLeveled, backendStdFormatter)
}

func Log() *Logger {
	return logger
}
