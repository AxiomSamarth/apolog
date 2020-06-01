package apolog

import (
	"log"
)

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
