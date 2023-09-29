package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

// Config holds data of configuration file
type Config struct {
	Server *Server `yaml:"server"`
	DB     *DB     `yaml:"db"`
	JWT    *JWT    `yaml:"jwt"`
}

// Load returns a *Config with all the configuration data
// needed to start the server
func Load(environment string) (*Config, error) {
	if environment == "" {
		return nil, fmt.Errorf("Environment not specified, got %s", environment)
	}

	//fileName := fmt.Sprintf("./config/%s.yml", environment)
	fileName := fmt.Sprintf("./../../config/%s.yml", environment)

	bytesYmlfile, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to read config file: %s, %s", fileName, err.Error())
	}

	cfg := &Config{}
	if err = yaml.Unmarshal(bytesYmlfile, cfg); err != nil {
		return nil, fmt.Errorf("unable to unmarshall configuration: %s", err.Error())
	}

	return cfg, nil
}
