package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

var config Configuration

func init() {
	log.Println("Loading config.yaml")
	content, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}
}

func Config() Configuration {
	return config
}
