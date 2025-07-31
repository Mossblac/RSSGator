package config

import (
	"os"
)

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homePath + configFileName, nil
}
