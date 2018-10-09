package config

import (
	"flag"
	"go-sample/utils/go-util/log"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	Port     int    `yaml:"port" json:"port"`
	Debug    bool   `yaml:"debug" json:"debug"`
	Timezone string `yaml:"timezone" json:"timezone"`
	Location *time.Location
}

var (
	AppConf AppConfig
)

func init() {
	flag.Parse()
}

func (c *AppConfig) Register() {
	ParseAppConfig()
}

//ParseAppConfig Parse application configs
func ParseAppConfig() {

	DefaultConfigurator.Load(`config/app`, &AppConf, func(config interface{}) {
		conf, _ := config.(*AppConfig)

		if conf.Timezone == `` {
			log.Fatal(`App timezone cannot be empty`)
		}
	})

	if os.Getenv(`APP_PORT`) != `` {
		if port, err := strconv.Atoi(os.Getenv(`APP_PORT`)); err == nil {
			AppConf.Port = port
		}

	}

	setDefaultTimeLocation(AppConf.Timezone)
}

//Set application default timezone
func setDefaultTimeLocation(timezone string) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatal(`Cannot load time location`, err)
	}

	AppConf.Location = location
}
