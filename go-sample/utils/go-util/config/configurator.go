package config

import "go-sample/utils/go-util/log"

var DefaultConfigurator Configurator

type Loader interface {
	Load(path string, i interface{}) error
}

type Configurator interface {
	Load(path string, i interface{}, validator func(config interface{}))
}

type defaultConfigurator struct {
	loader Loader
}

func NewConfigurator(loader Loader) *defaultConfigurator {
	return &defaultConfigurator{
		loader: loader,
	}
}

func (c *defaultConfigurator) Load(path string, i interface{}, validator func(config interface{})) {
	err := c.loader.Load(path, i)
	if err != nil {
		log.Fatal(`cannot load configuration `+path, err)
	}
	validator(i)
	log.Info(`configuration ` + path + ` loaded`)
}

func init() {
	DefaultConfigurator = NewConfigurator(NewYmlFileLoader())
}

type Config interface {
	Register()
}

type Configurations []Config

func LoadConfiguration(configs *Configurations) {
	for _, config := range *configs {
		config.Register()
	}
}
