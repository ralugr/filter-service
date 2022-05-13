package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	Warning    *log.Logger
	Info       *log.Logger
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

// init initializes loggers used all over the service
func init() {
	file, err := os.OpenFile(basepath+"/../../service-logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Failed to initialize logger: %v", err)
	}

	Info = createNewLogger("INFO: ", file)
	Warning = createNewLogger("WARNING: ", file)
}

// createNewLogger creates custom loggers
func createNewLogger(prefix string, file io.Writer) *log.Logger {
	newLogger := log.New(file, "FILTER: "+prefix, log.Ldate|log.Ltime|log.Lshortfile)
	mw := io.MultiWriter(os.Stdout, file)
	newLogger.SetOutput(mw)

	return newLogger
}
