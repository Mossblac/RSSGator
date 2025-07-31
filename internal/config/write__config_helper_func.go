package config

import (
	"encoding/json"
	"os"
)

func write(cfg Config) error {
	cfigPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(cfigPath, data, 0644)

}
