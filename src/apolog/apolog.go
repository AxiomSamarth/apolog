package apolog

import (
	"log"
)

// Info function implements Info level logging
func Info(data []bytes) {
	log.Print("INFO", data)
}

// Error function implements Error level logging
func Error(data []bytes) {
	log.Print("ERROR", data)
}

// Warn function implements Warn level logging
func Warn(data []bytes) {
	log.Print("WARNING", data)
}

// Debug function implements Debug level logging
func Debug(data []bytes) {
	log.Print("DEBUG", data)
}

// Trace function implements Trace level logging
func Trace(data []bytes) {
	log.Print("TRACE", data)
}

// Fatal function implements Fatal level logging
func Fatal(data []bytes) {
	log.Print("FATAL", data)
}
