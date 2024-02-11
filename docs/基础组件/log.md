
# 日志处理

在框架中有提供统一的日志实例，可以进行注入，也可以直接使用。可以直接复用 `tkit.Logger`。

或是在框架启动流程中注册自己自定义的 Logger 到 `tkit.Logger` 中就可以了。

## 默认组件

### 日志规范

在框架中，默认会把日志统一写入 `os.stdout` 与 `os.stderr`，日志格式为 JSON。

在日志类别上，默认分为：
- err：错误日志，业务应用和框架自身的报错、异常等信息均会写入该类别。
- access：访问日志，框架自身的所有访问和响应日志会写入该类别。一般不包含业务的日志，比较纯净。
- business：业务日志，业务应用打的各类业务信息，一般信息量较多较杂，常用于业务综合排查。

在日志存储上，建议按照日志类别和服务归类，如： /var/log/service/{日志类别}/{服务名称} 再辅以日志切割规则进行文件写入。

### 全局日志实例

```go
var Logger LoggerIface
```

### 日志接口方法

```go
type LoggerIface interface {
	// 业务日志
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})

	// 警告日志
	Warn(ctx context.Context, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})

	// 错误日志
	Error(ctx context.Context, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
}
```
