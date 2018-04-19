package log

import (
	"fmt"
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("example")
var formatFile = logging.MustStringFormatter(
	`%{time:01-02 15:04:05.000} %{level:.4s} %{shortpkg}.%{shortfunc} %{message}`,
)
var formatStd = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{level:.4s} %{shortpkg}.%{shortfunc} %{color:reset} %{message}`,
)

func Init(logfilename string, logFileLevel logging.Level) {
	logFile, err := os.OpenFile(logfilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	backendFile := logging.NewLogBackend(logFile, "", 0)
	backendStd := logging.NewLogBackend(os.Stdout, "", 0)

	backendFileFormatter := logging.NewBackendFormatter(backendFile, formatFile)
	backendStdFormatter := logging.NewBackendFormatter(backendStd, formatStd)

	backendFileLeveled := logging.AddModuleLevel(backendFileFormatter)
	backendFileLeveled.SetLevel(logFileLevel, "")

	logging.SetBackend(backendFileLeveled, backendStdFormatter)
}

func Log() *logging.Logger {
	return log
}
