package command

import (
	"fmt"
	"src/internal/pkg/handler"
)

type ApiRouteList struct {
}

func (s ApiRouteList) Name() string {
	return "api:route_list"
}

func (s ApiRouteList) Description() string {
	return "Show a list of all API http endpoints"
}

func (s ApiRouteList) Handle() error {
	fmt.Printf("\u001B[32mAll API endpoints:\u001B[0m\n")

	endpoints := handler.AppendApiByPath(ApiRoutes)
	for _, endpoint := range endpoints {
		fmt.Printf("%s\n", endpoint.Pattern)
	}

	return nil
}
