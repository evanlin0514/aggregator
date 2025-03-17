package main

import (
    "github.com/evanlin0514/aggregator/internal/config"
    "fmt"
)

func main() {
    file, err := config.Read()
    if err != nil {
        fmt.Println(err)
    }
    
    err = file.SetUser("Evan")
    if err != nil {
        fmt.Println(err)
    }
    

    newFile, err := config.Read()
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(newFile)
}
