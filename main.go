package main

import (
	"fmt"
	"os"

	"github.com/evanlin0514/aggregator/internal/config"
	"github.com/evanlin0514/aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
    db *database.Queries
    cfg *config.Config
}

func main() {
    file, err := config.Read()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    state := config.State{
        Pointer: &file,
    }

    if len(os.Args) < 3 {
        fmt.Println("invlid input")
		os.Exit(1)
	}

    input := os.Args
    cmd := config.Command{
        Name: input[1],
        Args: input[2:],
    }

    cmds := config.Commands{
        Handlers: make(map[string]func(*config.State, config.Command) error),
    }
    cmds.Register(cmd.Name, config.HandlerLogin)

    err = cmds.Run(&state, cmd)
    if err != nil {
        fmt.Printf("error when runing %v\n", err)
        os.Exit(1)
    }

}
