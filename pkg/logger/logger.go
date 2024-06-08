package logger

import (
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	logrus "github.com/sirupsen/logrus"
)

// Log - logrus logging.
//
//revive:disable:import-shadowing
var (
	Log       *logrus.Logger
	LogStdErr *logrus.Logger
)

func Init() {
	logStdOut := logrus.New()
	logStdOut.SetOutput(os.Stdout)
	logStdOut.SetFormatter(&nested.Formatter{
		TimestampFormat: "Jan 02 03:04:05.000 PM",
		HideKeys:        false,
		FieldsOrder:     []string{"component", "category"},
	})
	Log = logStdOut

	logStdErr := logrus.New()
	logStdErr.SetOutput(os.Stderr)
	logStdErr.SetFormatter(&nested.Formatter{
		TimestampFormat: "Jan 02 03:04:05.000 PM",
		HideKeys:        false,
		FieldsOrder:     []string{"component", "category"},
	})
	LogStdErr = logStdErr
}
