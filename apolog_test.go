package apolog_test

import (
	"testing"

	apolog "github.com/axiomsamarth/apolog/apolog"
)

// TestInfo function tests Info level of logging
func TestInfo(t *testing.T) {
	var data = []byte(`Log this!`)
	apolog.Info(data)
}

// TestError function tests Info level of logging
func TestError(t *testing.T) {
	var data = []byte(`Log this!`)
	apolog.Error(data)
}

// TestWarn function tests Info level of logging
func TestWarn(t *testing.T) {
	var data = []byte(`Log this!`)
	apolog.Warn(data)
}

// TestDebug function tests Info level of logging
func TestDebug(t *testing.T) {
	var data = []byte(`Log this!`)
	apolog.Debug(data)
}

// TestTrace function tests Info level of logging
func TestTrace(t *testing.T) {
	var data = []byte(`Log this!`)
	apolog.Trace(data)
}
