package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (path string, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configFilePath := filepath.Join(homeDir, configFileName)
	return configFilePath, nil
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func write(cfg *Config) error {
	configFilePath, _ := getConfigFilePath()
	configData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(configFilePath, configData, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user
	err := write(cfg)
	if err != nil {
		return err
	}
	return nil
}
