package command

import (
    "fmt"
    "src/config"
    "src/internal/pkg/handler"
)

type ApiList struct {
}

func (s ApiList) Name() string {
    return "api:list"
}

func (s ApiList) Description() string {
    return "Show a list of all API http endpoints"
}

func (s ApiList) Handle() error {
    fmt.Printf("\u001B[32mAll API endpoints:\u001B[0m\n")

    endpoints := handler.AppendApiByPath(ApiRoutes)
    for _, endpoint := range endpoints {
        fmt.Printf("%s\n", endpoint.Pattern)
    }

    return nil
}

func (s ApiList) getListenAddr() string {
    return fmt.Sprintf("%s:%d", config.AppServe.Host, config.AppServe.Port)
}
