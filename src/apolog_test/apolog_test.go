package apolog_test

import (
	"testing"

	apolog "github.com/axiomsamarth/apolog"
)

// TestInfo function tests Info level of logging
func TestInfo(t *testing.T) {
	apolog.Info(`Log this!`)
}
