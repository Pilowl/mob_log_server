package config

import (
	"encoding/json"
	"os"
)

type (
	Config struct {
		Port string
		DB   Database
	}

	Database struct {
		Username string
		Password string
		Port     string
		Name     string
	}
)

var configuration *Config

func GetConfig() *Config {
	return configuration
}

func Init(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic("Cannot parse configuration file. Check 'config' directory.")
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		panic("Cannot parse configuration file. Check 'config' directory JSON config files.")
	}
}
