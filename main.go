package main

import (
	"database/sql"
	"log"
	"os"
	"github.com/evanlin0514/aggregator/internal/config"
	"github.com/evanlin0514/aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct{
    db *database.Queries
    pointer *config.Config
}

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

func main() {
    file, err := config.Read()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    db, err := sql.Open("postgres", file.DbUrl)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err )
    }

    dbQueries := database.New(db)

    programState := &state{
        pointer: &file,
        db: dbQueries,
    }


    input := os.Args
    cmd := command{
        name: input[1],
        args: input[2:],
    }

    cmds := commands{
        Handlers: make(map[string]func(*state, command) error),
    }

    cmds.register("login", handlerLogin)
    cmds.register("register", handlerRegister)
    cmds.register("reset", handlerReset)
    cmds.register("users", handlerUsers)

    err = cmds.run(programState, cmd)
    if err != nil {
        log.Fatalf("Error runing command: %v", err)
    }
    

}
