package helper

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

func NewLogger() *logrus.Logger {
	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.InfoLevel,
		// Level:        convertLogLevel(DefaultLogLevel),
		ReportCaller: true,
		ExitFunc:     os.Exit,
		Hooks:        logrus.LevelHooks{},
	}

	return log
}

func NewTracerLogger() *logrus.Logger {
	l := NewLogger()
	l.SetFormatter(NewDefaultFormatter())
	return l
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
