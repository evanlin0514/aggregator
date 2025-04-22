package main

import(
	"fmt"
	"context"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("invalid input, bad arg")
	}

	id, err := s.db.GetId(context.Background(), s.pointer.CurrentUser)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), id)
	if err != nil {
		return fmt.Errorf("error getting feeds user follow: %v", err)
	}

	for _, feed := range(feeds){
		fmt.Println(feed.FeedName)
	}
	return nil
}


