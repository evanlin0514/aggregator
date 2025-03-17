package config

import (
	"encoding/json"
	"path/filepath"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUser string  `json:"current_user_name"`
}

func getFilePath(file string) (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home dir: %v", err)
	}

	return filepath.Join(path,file), err
}

func Read() (Config, error) {
	var c Config
	path, err := getFilePath(configFileName)
	if err != nil {
		return c, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err){
		return c, fmt.Errorf("config file does not exist: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return c, fmt.Errorf("error reading file: %v", err)
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		return c, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	return c, nil
}

func write(c Config) error {
	djson, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshaling struct: %v", err)
	}

	path, err := getFilePath(configFileName)
	if err != nil {
		return err
	}
	
	err = os.WriteFile(path, djson, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}
	return nil
}

func (c *Config) SetUser (name string) error {
	c.CurrentUser = name
	return write(*c)
}




