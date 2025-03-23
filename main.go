package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/evanlin0514/aggregator/internal/config"
	"github.com/evanlin0514/aggregator/internal/database"
	_ "github.com/lib/pq"
)


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

    state := &config.State{
        Pointer: &file,
        Db: dbQueries,
    }


    if len(os.Args) < 3 {
        log.Fatal("invalid input")
	}

    input := os.Args
    cmd := config.Command{
        Name: input[1],
        Args: input[2:],
    }

    cmds := config.Commands{
        Handlers: make(map[string]func(*config.State, config.Command) error),
    }

    cmds.Register("login", config.HandlerLogin)
    cmds.Register("register", config.HandlerRegister)

    err = cmds.Run(state, cmd)
    if err != nil {
        log.Fatalf("Error runing command: %v", err)
    }
    

}
