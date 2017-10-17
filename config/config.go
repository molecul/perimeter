package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type PerimeterConfig struct {
	ConfigPath string
	Perimeter  struct {
		Consul struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		}
		Server struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		}
		Themis struct {
			URL string `yaml:"url"`
		}
	}
}

func (c *PerimeterConfig) Read() {
	res, err := ioutil.ReadFile(c.ConfigPath)
	if err != nil {
		panic(err)
	}

	yaml.Unmarshal([]byte(res), &c)
}
