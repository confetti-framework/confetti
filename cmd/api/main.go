package main

import (
    "fmt"
    "os"
    "src/cmd/api/command"
)

type Command interface {
    Name() string
    Description() string
    Handle() error
}

var commands = []Command{
    command.ApiList{},
    command.AppServe{},
    // Here you can add your own custom commands
}

func main() {
    // First argument
    if len(os.Args) <= 1 {
        fmt.Printf("All commands:\n\n")
        for _, cmd := range commands {
            fmt.Printf("%s\t%s\n", cmd.Name(), cmd.Description())
        }
        return
    }
    // Get command
    cmd, err := getCommandByName(os.Args[1])
    if err != nil {
        fmt.Printf(err.Error())
        os.Exit(1)
    }
    // Handle command
    err = cmd.Handle()
    if err != nil {
        fmt.Printf("%s\n", err.Error())
        os.Exit(1)
    }
}

func getCommandByName(name string) (Command, error) {
    for _, cmd := range commands {
        if cmd.Name() == name {
            return cmd, nil
        }
    }
    return nil, fmt.Errorf("command not found\n")
}
