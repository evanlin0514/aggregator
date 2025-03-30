package main

import(
	"context"
	"fmt"
	"strings"
	"github.com/evanlin0514/aggregator/internal/database"
)

func handlerFeed (s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("invalid input")
	}

	user_id, err := s.db.GetId(context.Background(), s.pointer.CurrentUser)
	if err != nil{
		return fmt.Errorf("error getting user id: %v", err)
	}

	params := database.AddfeedParams{
		Name: cmd.args[0],
		Url: cmd.args[1],
		UserID: user_id,
	}

	feed, err := s.db.Addfeed(context.Background(), params)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key"){
			return fmt.Errorf("post already exists")
		}
		return fmt.Errorf("error adding feed: %v", err)
	}

	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("ID:		        %s\n", feed.ID)
	fmt.Printf("Created at:		%s\n", feed.CreatedAt)
	fmt.Printf("Updated at:		%s\n", feed.UpdatedAt)
	fmt.Printf("Name:			  %s\n", feed.Name)
	fmt.Printf("Url:			   %s\n", feed.Url)
	fmt.Printf("User ID:		   %s\n", feed.UserID)

}