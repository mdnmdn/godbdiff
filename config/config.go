package config

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Source  Connection `yaml:"source,omitempty"`
	Target  Connection `yaml:"target,omitempty"`
	Schemas []string   `yaml:"schemas"`
}

type Connection struct {
	ConnectionString string `yaml:"connectionString,omitempty"`
}

func LoadConfigFromFile(filename string) (Config, error) {
	var config Config
	configPath, err := filepath.Abs(filename)
	if err != nil {
		return config, err
	}

	yaml, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	return LoadConfig(yaml)
}

func LoadConfig(configData []byte) (Config, error) {
	var config Config
	err := yaml.Unmarshal(configData, &config)

	return config, err
}
