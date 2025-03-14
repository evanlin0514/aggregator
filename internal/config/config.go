package config

import (
	"encoding/json"
	"path/filepath"
	"fmt"
	"os"
)

const configFileName = ".gotorconfig.json"

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUser string  `json:"current_user_name"`
}

func getFilePath(file string) (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home dir: %v", err)
	}

	if _, err := os.Stat(filepath.Join(path,file)); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %v", filepath.Join(path, file))
	}

	return filepath.Join(path,file), err
}

func Read(file string) (Config, error) {
	var c Config
	path, err := getFilePath(file)
	if err != nil {
		return c, err
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

func write(file string, c Config) error {
	djson, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshaling struct: %v", err)
	}
	err = os.WriteFile(file, djson, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}
	return nil
}

func (c *Config) SetUser (name string) error {
	c.CurrentUser = name
	path, err := getFilePath(configFileName)
	if err != nil {
		return err
	}
	err = write(path, *c)
	if err != nil {
		return err
	}
	return nil
}




