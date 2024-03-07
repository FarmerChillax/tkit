package helper

import (
	"os"

	"github.com/FarmerChillax/tkit/config"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

func NewLogger(loggerConf *config.LoggerConfig) (*logrus.Logger, error) {
	log := &logrus.Logger{
		Out: os.Stdout,
		// Level: logrus.InfoLevel,
		Level:        logrus.Level(loggerConf.Level),
		ReportCaller: loggerConf.ReportCaller,
		ExitFunc:     os.Exit,
		Hooks:        logrus.LevelHooks{},
	}

	return log, nil
}

func NewTracerLogger(loggerConf *config.LoggerConfig) (*logrus.Logger, error) {
	l, err := NewLogger(loggerConf)
	if err != nil {
		return nil, err
	}
	l.SetFormatter(NewDefaultFormatter())
	return l, nil
}

type Formatter struct {
	formatter logrus.Formatter
}

func NewDefaultFormatter() logrus.Formatter {
	return &Formatter{
		formatter: &logrus.JSONFormatter{},
	}
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 不输出 file 字段
	delete(entry.Data, "file")

	span := trace.SpanFromContext(entry.Context)
	entry.Data["trace_id"] = span.SpanContext().TraceID().String()
	entry.Data["span_id"] = span.SpanContext().SpanID().String()
	return f.formatter.Format(entry)
}
