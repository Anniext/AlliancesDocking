package config

type ConfigsManager interface {
	Read()
	Write()
	Detection()
}

type OptionConfig struct {
	Option string
}

var (
	configs      ConfigsManager
	logs         ConfigsManager
	jsonConfBody JsonConfBody
	jsonLogBody  JsonLogBody
)

func init() {
	configs = &OptionConfig{Option: NewConfig()}
	logs = &OptionConfig{Option: NewLog()}
}

func Configinit() {
	configs.Read()
	logs.Write()
	logs.Detection()
}

func GetConfig() *JsonConfBody {
	return &jsonConfBody
}

func GetLog() *JsonLogBody {
	return &jsonLogBody
}
