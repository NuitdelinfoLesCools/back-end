package main

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// Config .
type Config struct {
	Database struct {
		Host             string  `yaml:"host"`
		User             string  `yaml:"user"`
		Password         string  `yaml:"password`
		Name             string  `yaml:"name"`
		Port             int     `yaml:"port"`
		SslMode          int8    `yaml:"ssl_model"`
		MaxIdleConns     int     `yaml:"max_idle_conns"`
		MaxOpenConns     int     `yaml:"max_open_conns"`
		MaxLifetimeConns float64 `yaml:"max_lifetime_conns"`
	} `yaml":database"`
	Port        int    `yaml:"port"`
	StaticFiles string `yaml:"static_files"`
	StaticURL   string `yaml:"static_url"`
	ApiKeys     struct {
		OpenWeatherApi string `yaml:"open_weather_api"`
	} `yaml:"api_keys"`
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
