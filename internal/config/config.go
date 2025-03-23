package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/evanlin0514/aggregator/internal/database"
	"github.com/google/uuid"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUser string  `json:"current_user_name"`
}

type State struct {
	Db *database.Queries
	Pointer *Config	
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handlers map[string]func(*State, Command)error 
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

func HandlerLogin(s *State, cmd Command) error{
	_, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil{
		return fmt.Errorf("no user found: %v", err)
	}
	s.Pointer.SetUser(cmd.Args[0])
	fmt.Printf("successfully swtich to user: %v \n", cmd.Args[0])
	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Handlers[name] = f
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	params := database.CreateUserParams{
		ID: uuid.New(),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		Name: username,
	}

	newUser, err := s.Db.CreateUser(context.Background(), params)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key"){
			return fmt.Errorf("user already exists")
		}
		return fmt.Errorf("error creating user: %v", err)
	}

	if err := s.Pointer.SetUser(username); err != nil{
		return err
	}

	fmt.Printf("User created: %v\n", newUser.Name)
	return nil
}

func (c *Commands) Run (s *State, cmd Command) error{
	f, ok := c.Handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("handler not exist")
	}

	err := f(s, cmd)
	if err != nil {
		return err
	}

	return nil
}


