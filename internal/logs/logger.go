package logs

import (
	"log"
	"os"
	"path/filepath"
)

var (
	logFile *os.File
	Logger  *log.Logger
)

// Init инициализирует глобальный логгер, пишет в ~/.vpn_auto/log.txt
func Init() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	logPath := filepath.Join(home, ".vpn_auto", "log.txt")

	err = os.MkdirAll(filepath.Dir(logPath), 0755)
	if err != nil {
		return err
	}

	logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	Logger = log.New(logFile, "", log.LstdFlags)
	return nil
}

// Close закрывает лог-файл
func Close() {
	if logFile != nil {
		logFile.Close()
	}
}
