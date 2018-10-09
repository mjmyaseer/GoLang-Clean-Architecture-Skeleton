package config

import (
	"go-sample/utils/go-util/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type JsonFileLoader struct{}

func NewJsonFileLoader() Loader {
	return new(JsonFileLoader)
}

func (l *JsonFileLoader) Load(path string, i interface{}) error {

	file := path + `.json`
	byt, err := ioutil.ReadFile(file)
	if err != nil {
		log.Error(`cannot load file `, err)
		return err
	}

	log.Info(`Config ` + file + ` loaded`)

	return yaml.Unmarshal(byt, i)
}
