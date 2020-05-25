package internal

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
}

type config struct {
	DB db `yaml:"db"`
}

var Config *config

func init() {
	path, _ := os.Getwd()
	file, _ := ioutil.ReadFile(path + "/config.yaml")
	yaml.Unmarshal(file, &Config)
}
