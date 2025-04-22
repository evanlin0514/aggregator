package main

import (
	"context"
	"fmt"

	"github.com/evanlin0514/aggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("invalid input, command arg should be one")
	}

	user_id, err := s.db.GetId(context.Background(), s.pointer.CurrentUser)
	if err != nil{
		return fmt.Errorf("error getting user id: %v", err)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil{
		return fmt.Errorf("invalid feed url: %v", err)
	}

	param := database.CreateFeedFollowParams{
		UserID: user_id,
		FeedID: feed.ID,
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), param)
	if err != nil {
		return fmt.Errorf("error create feed follow: %v", err)
	}

	fmt.Println("feed name: ", follow.FeedName)
	fmt.Println("current user: ", follow.UserName)
	return nil
}

