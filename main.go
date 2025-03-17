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

    input := os.Args[1:]
    loginCmd := config.Command{
        Name: "login",
        Args: input,
    }

    cmds := config.Commands{
        Handlers: make(map[string]func(*config.State, config.Command) error),
    }
    cmds.Register(loginCmd.Name, config.HandlerLogin)

    err = cmds.Run(&state, loginCmd)
    if err != nil {
        fmt.Printf("error when runing %v\n", err)
        os.Exit(1)
    }
}
