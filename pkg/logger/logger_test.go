package logger

import (
	"testing"

	nested "github.com/antonfisher/nested-logrus-formatter"
	logrus "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLoggerInitialization(t *testing.T) {
	// Save the original logger to restore it after the test
	originalLogger := LogStdErr

	// Create a new logger for testing purposes
	testLogger := logrus.New()
	testLogger.SetFormatter(&nested.Formatter{
		TimestampFormat: "Jan 02 03:04:05.000 PM",
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	})

	// Replace the global logger with the test logger
	LogStdErr = testLogger

	assert.Equal(t, testLogger.Formatter, LogStdErr.Formatter)
	assert.Equal(t, testLogger.Out, LogStdErr.Out)

	// Restore the original logger to avoid interfering with other tests
	LogStdErr = originalLogger
}
