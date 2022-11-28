package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Api struct {
		Key    string `yaml:"key"`
		Secret string `yaml:"secret"`
	} `yaml:"api"`
}

func GetConfig(path string) *Config {
	var conf Config

	confContent, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	errYaml := yaml.Unmarshal(confContent, &conf)
	if errYaml != nil {
		panic(err.Error())
	}

	return &conf
}
