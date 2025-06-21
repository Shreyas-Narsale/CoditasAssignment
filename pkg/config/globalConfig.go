package config

type GlobalConfig struct {
}

var globalConf = &GlobalConfig{}

func GetGlobalConfig() *GlobalConfig {
	return globalConf
}

func LoadGlobalConfig() {

}
