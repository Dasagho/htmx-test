// Archivo: logging/logger.go
package logging

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	debugLogger *log.Logger
	errorLogger *log.Logger
	infoLogger  *log.Logger
	logDir      string
)

func init() {
	logDir = "log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, 0755)
	}

	debugLogger = initLogger("debug.log")
	errorLogger = initLogger("error.log")
	infoLogger = initLogger("info.log")
}

func initLogger(filename string) *log.Logger {
	filePath := filepath.Join(logDir, filename)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error al abrir el archivo de log: %v", err)
	}

	// Formato personalizado para el logger
	flag := log.Lshortfile
	logger := log.New(file, "", flag)
	logger.SetFlags(0)   // Desactivar flags predeterminadas
	logger.SetPrefix("") // Desactivar prefijo predeterminado

	return logger
}

func logFormatted(logger *log.Logger, level, message string) {
	logger.Printf("[%s] %s: %s", time.Now().Format("2006-01-02 15:04:05"), level, message)
}

func Debug(message string) {
	logFormatted(debugLogger, "DEBUG", message)
}

func Error(message string) {
	logFormatted(errorLogger, "ERROR", message)
}

func Info(message string) {
	logFormatted(infoLogger, "INFO", message)
}
