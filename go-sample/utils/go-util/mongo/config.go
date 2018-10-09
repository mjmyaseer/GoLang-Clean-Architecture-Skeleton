package mongo

import (
	//"go-sample/utils/go-util/log"
	//"gopkg.in/yaml.v2"
	//"io/ioutil"
	"go-sample/utils/go-util/config"
)

type Config struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Database string `yaml:"database" json:"database"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
	Auth     bool   `yaml:"auth" json:"auth"`
	AuthDb   string `yaml:"auth_db" json:"auth_db"`
}

var Conf Config

func (Config) Register() {
	config.DefaultConfigurator.Load(`config/mongo`, &Conf, func(config interface{}) {})
}

func (Config) Message() string {
	return `Mongo config loaded`
}

func (Config) Validate() {}
