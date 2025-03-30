package main

import (
	"io"
	"fmt"
	"html"
	"bytes"
	"context"
	"net/http"
	"encoding/xml"
)


type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func catchEscape(rss *RSSFeed) {
	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)
	for _, item := range rss.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){
	rss := &RSSFeed{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, &bytes.Reader{})
	if err != nil {
		return rss, fmt.Errorf("error making new request: %v", err)
	}

	req.Header.Add("User-Agent", "gator")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return rss, fmt.Errorf("error when GET: %v", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return rss, fmt.Errorf("error reading Body: %v", err)
	}
	if err := xml.Unmarshal(body, rss); err != nil{
		return rss, fmt.Errorf("error unmarshaling Body: %v", err)
	}

	catchEscape(rss)
	return rss, nil
}

