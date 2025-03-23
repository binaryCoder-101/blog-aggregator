package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string, filePath string) error {
	c.CurrentUserName = userName

	jsonData, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error parsing config data struct")
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file")
	}

	return nil
}
