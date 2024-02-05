package helper

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

func NewLogger(formatter logrus.Formatter) *logrus.Logger {
	if formatter == nil {
		formatter = NewDefaultFormatter()
	}

	log := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: formatter,
		Level:     logrus.InfoLevel,
		// Level:        convertLogLevel(DefaultLogLevel),
		ReportCaller: true,
		ExitFunc:     os.Exit,
		Hooks:        logrus.LevelHooks{},
	}

	return log
}

type Formatter struct {
	formatter logrus.JSONFormatter
}

func NewDefaultFormatter() logrus.Formatter {
	return &Formatter{
		formatter: logrus.JSONFormatter{},
	}
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	span := trace.SpanFromContext(entry.Context)
	entry.Data["trace_id"] = span.SpanContext().TraceID().String()
	entry.Data["span_id"] = span.SpanContext().SpanID().String()
	return f.formatter.Format(entry)
}
