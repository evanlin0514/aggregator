package main

import(
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("invalid input")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}

	for _, feed := range feeds {
		name, err := s.db.GetName(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting username: %v", err)
		}

		fmt.Println("------------------POST------------------------")
		fmt.Printf("Name:			   %s\n", feed.Name)
		fmt.Printf("Url:			   %s\n", feed.Url)
		fmt.Printf("User ID:		   %s\n", name)
		fmt.Println("----------------------------------------------")
	}
	return nil
}
