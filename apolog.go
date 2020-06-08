package apolog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Apolog sets the context for filesystem logging
type Apolog struct {
	Filepath string
}

// SetContext will set the context of the log files
func (context *Apolog) SetContext() {
	_, err := filepath.Abs(context.Filepath)
	if err != nil {
		fmt.Println("Error locating the log file")
	}
	logfile, err := os.OpenFile(context.Filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(logfile)
}

// Info function implements Info level logging
func Info(data []byte) {
	log.Print("INFO: ", string(data))
}

// Error function implements Error level logging
func Error(data []byte) {
	log.Print("ERROR: ", string(data))
}

// Warn function implements Warn level logging
func Warn(data []byte) {
	log.Print("WARNING: ", string(data))
}

// Debug function implements Debug level logging
func Debug(data []byte) {
	log.Print("DEBUG: ", string(data))
}

// Trace function implements Trace level logging
func Trace(data []byte) {
	log.Print("TRACE: ", string(data))
}

// Fatal function implements Fatal level logging
func Fatal(data []byte) {
	log.Fatal("FATAL: ", string(data))
}
