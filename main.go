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
    cmds.register("agg", handlerAgg)
    cmds.register("addfeed", handlerFeed)

    err = cmds.run(programState, cmd)
    if err != nil {
        log.Fatalf("Error runing command: %v", err)
    }
    

}
