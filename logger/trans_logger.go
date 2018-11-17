package logger

import (
	"bufio"
	"io"
	"log"
	"os"
	"path"

	"github.com/ggvishnu29/horlix/model"
)

var transLog *os.File
var tLogger *log.Logger
var tWriter *bufio.Writer
var tRecoveryFlag bool
var transLogPath string
var transLock model.Lock

//InitTransLogger creates a transaction log
func InitTransLogger(transLogDir string) error {
	var err error
	transLogPath = path.Join(transLogDir, "transaction.log")
	transLog, err = os.Create(transLogPath)
	if err != nil {
		return err
	}
	tWriter = bufio.NewWriter(transLog)
	return nil
}

//SetTransLogRecoveryFlag sets the Recovery flag
func SetTransLogRecoveryFlag() {
	tRecoveryFlag = true
}

//UnSetTransLogRecoveryFlag unsets the Recovery flag
func UnSetTransLogRecoveryFlag() {
	tRecoveryFlag = false
}

//CopyTruncateTransLogToFile does the copy truncate
func CopyTruncateTransLogToFile(destFile string) error {
	LogInfo("acquiring trans log lock")
	transLock.Lock()
	defer transLock.UnLock()
	LogInfo("trans lock acquired")
	in, err := os.Open(transLogPath)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	if err := truncateTransLog(); err != nil {
		return err
	}
	return nil
}

func truncateTransLog() error {
	if err := transLog.Truncate(0); err != nil {
		return err
	}
	if _, err := transLog.Seek(0, 0); err != nil {
		return err
	}
	return nil
}

//LogTransactions writes the  transaction message to the log
func LogTransactions(logs [][]byte) error {
	transLock.Lock()
	defer transLock.UnLock()
	for _, log := range logs {
		if _, err := tWriter.WriteString(string(log[:]) + "\n"); err != nil {
			return err
		}
	}
	if err := tWriter.Flush(); err != nil {
		return err
	}
	return nil
}
