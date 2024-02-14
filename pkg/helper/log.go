package helper

import (
	"fmt"
	"io"
	"os"

	"github.com/FarmerChillax/tkit/config"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

func NewLogger(loggerConf *config.LoggerConfig) (*logrus.Logger, error) {
	f, err := os.OpenFile(loggerConf.OutputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	multiWriter := io.MultiWriter(os.Stdout, f)
	log := &logrus.Logger{
		Out: multiWriter,
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
	fmt.Println("Formatter.Format: ", entry.Data)
	// 不输出 file 字段
	delete(entry.Data, "file")

	span := trace.SpanFromContext(entry.Context)
	entry.Data["trace_id"] = span.SpanContext().TraceID().String()
	entry.Data["span_id"] = span.SpanContext().SpanID().String()
	return f.formatter.Format(entry)
}
