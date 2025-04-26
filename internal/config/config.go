package config

import (
	"encoding/json"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Read reads the config file from the user's home directory
// and unmarshals it into a Config struct
func Read() (Config, error) {

	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

// SetUser sets the current user in the config file
// and writes the updated config back to the file
func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(*c)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// GetConfigFilePath returns the full path to the config file
func getConfigFilePath() (string, error) {

	fullPath := ""

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fullPath, err
	}

	fullPath = homeDir + configFileName

	return fullPath, nil
}

// func write() {

// }
