package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("error getting home directory")
	}

	filePath := homeDir + "/.gatorconfig.json"

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading json data")
	}

	config := Config{}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding json data")
	}

	return config, nil
}
