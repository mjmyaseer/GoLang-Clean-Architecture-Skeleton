package configurations

import "go-sample/utils/go-util/config"

type UtilConfig struct {
	Expiry        int64   `yaml:"backend_expiry_time" json:"backend_expiry_time"`
	NumOfThreads  int     `yaml:"num_of_threads" json:"num_of_threads"`
	Timeout       float64 `yaml:"expiry_timeout" json:"expiry_timeout"`
	CleanInterval float64 `yaml:"clean_interval" json:"clean_interval"`
}

var UtilConf UtilConfig

func (s *UtilConfig) Register() {
	config.DefaultConfigurator.Load(`config/util`, &UtilConf, func(config interface{}) {
	})
}
