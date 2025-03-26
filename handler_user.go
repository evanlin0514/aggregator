package main

import (
	"context"
	"fmt"
	"time"
	"strings"
	"github.com/google/uuid"
	"github.com/evanlin0514/aggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error{
	_, err := s.db.CheckUser(context.Background(), cmd.args[0])
	if err != nil{
		return fmt.Errorf("no user found: %v", err)
	}
	s.pointer.SetUser(cmd.args[0])
	fmt.Printf("successfully swtich to user: %v \n", cmd.args[0])
	return nil
}


func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}
	username := cmd.args[0]

	params := database.CreateUserParams{
		ID: uuid.New(),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		Name: username,
	}

	newUser, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key"){
			return fmt.Errorf("user already exists")
		}
		return fmt.Errorf("error creating user: %v", err)
	}

	if err := s.pointer.SetUser(username); err != nil{
		return err
	}

	fmt.Printf("User created: %v\n", newUser.Name)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) == 1 {
		return fmt.Errorf("invlid input")
	}

	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("error reset table: %v", err)
	}

	fmt.Println("table is reset")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) == 1 {
		return fmt.Errorf("invlid input")
	}	
	
	users, err := s.db.GetUser(context.Background())
	if err != nil {
		return fmt.Errorf("error retriving users: %v", err)
	}

	if len(users) > 0 {
		for _, user := range users{
			if user == s.pointer.CurrentUser {
				fmt.Printf("* %v (current)\n", user)
			} else {
				fmt.Printf("* %v \n", user)
			}
		}
	} else {
		fmt.Println("the table is empty")
	}

	return nil
}