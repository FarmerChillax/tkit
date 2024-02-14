package config

import sdktrace "go.opentelemetry.io/otel/sdk/trace"

type OtelConfig struct {
	Exporter sdktrace.SpanExporter
}

// 获取 otel 配置
// todo...
func getOtelConfigFromEnv() *OtelConfig {
	otelConf := OtelConfig{}
	return &otelConf
}
