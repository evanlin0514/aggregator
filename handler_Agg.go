package main

import (
	"context"
	"fmt"
)


func handlerAgg (s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("invlid input")
	}

	rss, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
    if err != nil {
        return fmt.Errorf("error fetching feed: %v", err)
    }

	fmt.Println(rss)
	return nil
}