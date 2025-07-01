package tkit

type ApplicationOptionType int32

const (
	// 应用配置中心
	ApplicationConfigerType ApplicationOptionType = 1
)

type OptionIface interface{}

// OptionDisable used to close app option
type OptionDisable struct {
	OptionIface
}
