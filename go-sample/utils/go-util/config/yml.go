package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type YmlFileLoader struct{}

func NewYmlFileLoader() Loader {
	return new(YmlFileLoader)
}

func (YmlFileLoader) Load(path string, i interface{}) error {

	file := path + `.yaml`
	byt, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return yaml.UnmarshalStrict(byt, i)

}
