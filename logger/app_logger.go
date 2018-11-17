package logger

import (
	"bufio"
	"log"
	"os"
	"path"
)

var aLogger *log.Logger
var appLog *os.File

//InitAppLogger creates the application log file.
func InitAppLogger(appLogDir string) error {
	var err error
	appLog, err = os.OpenFile(path.Join(appLogDir, "horlix.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	buf := bufio.NewWriterSize(appLog, 10)
	aLogger = log.New(buf, "", log.Lshortfile)
	return nil
}

//LogInfo logs a info message
func LogInfo(msg string) {
	aLogger.Println(msg)
}

//LogInfof logs a formatted info message
func LogInfof(format string, v ...interface{}) {
	aLogger.Printf(format, v...)
}

//LogFatal logs a fatal message
func LogFatal(msg string) {
	aLogger.Fatalln(msg)
}

//LogFatalf logs a formatted fatal message
func LogFatalf(format string, v ...interface{}) {
	aLogger.Fatalf(format, v...)
}
