package main

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// Config .
type Config struct {
	Database struct {
		Host             string
		User             string
		Password         string
		Name             string
		Port             int
		SslMode          int8
		MaxIdleConns     int
		MaxOpenConns     int
		MaxLifetimeConns float64
	} `yaml":database"`
	Port        int    `yaml:"port"`
	StaticFiles string `yaml:"static_files"`
	StaticURL   string `yaml:"static_url"`
}

func loadConfig(path string) (*Config, error) {
	var config Config
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
