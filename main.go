package main

import (
	"fmt"
	"os"

	"github.com/evanlin0514/aggregator/internal/config"
)

func main() {
    file, err := config.Read()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    state := config.State{
        Pointer: &file,
    }

    cmd := config.Command{
        Name: os.Args[1],
        Args: os.Args[1:],
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
