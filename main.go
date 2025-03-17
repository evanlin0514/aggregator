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
    }
    state := config.State{
        Pointer: &file,
    }
    
    var cmds config.Commands
    cmds.Register("login", config.HandlerLogin(&state, os.Args))

}
