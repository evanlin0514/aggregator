package main

import (
    "github.com/evanlin0514/aggregator/internal/config"
    "fmt"
)

func main() {
    file := config.Config{} 
    err := config.Read(".gatorconfig.json", &file)
    if err != nil {
        fmt.Println(err)
    }
    
    err = file.SetUser("Evan")
    if err != nil {
        fmt.Println(err)
    }

    err = config.Read("gatorconfig.json", &file)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(file)
}
