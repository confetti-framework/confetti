package main

import (
    "fmt"
    "os"
    "src/app/console"
)

type Command interface {
    Name() string
    Description() string
    Handle() error
}

var commands = []Command{
    // Here you can add your own custom commands
    console.AppServe{},
}

func main() {
    // First argument 
    if len(os.Args) <= 1 {
        fmt.Printf("All commands:\n\n")
        for _, command := range commands {
            fmt.Printf("%s\t%s\n", command.Name(), command.Description())
        }
        return
    }
    // Get command
    command, err := getCommandByName(os.Args[1])
    if err != nil {
        fmt.Printf(err.Error())
        os.Exit(1)
    }
    // Handle command
    err = command.Handle()
    if err != nil {
        fmt.Printf("%s\n", err.Error())
        os.Exit(1)
    }
}

func getCommandByName(name string) (Command, error) {
    for _, command := range commands {
        if command.Name() == name {
            return command, nil
        }
    }
    return nil, fmt.Errorf("command not found\n")
}
