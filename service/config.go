package service

import (
	"golang-demo/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

//Load config
func GetConfig() (*model.Config, error) {
	conf, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	d := &model.Config{}
	err = yaml.Unmarshal(conf, d)
	if err != nil {
		return nil, err
	}
	return d, nil
}
