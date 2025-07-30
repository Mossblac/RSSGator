package config

import (
	"encoding/json"
	"os"
)

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homePath + configFileName, nil
}

func Read() (Config, error) {
	var config Config

	cfig, err := getConfigFilePath()
	if err != nil {
		return Config{}, nil
	}
	data, err := os.ReadFile(cfig)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil

}
