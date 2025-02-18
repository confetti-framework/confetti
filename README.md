<a href="https://github.com/confetti-framework/confetti/blob/main/test/README.md"><img align="right" src="https://img.shields.io/badge/average_coverage-87%25-yellowgreen"></a><br>
<a href="https://goreportcard.com/report/github.com/confetti-framework/confetti"><img align="right" src="https://goreportcard.com/badge/github.com/confetti-framework/confetti"></a><br>
<a href="https://pkg.go.dev/github.com/confetti-framework/confetti"><img align="right" src="https://godoc.org/confetti-framework/confetti?status.svg"></a>
<br>
<a href="https://confetti-framework.github.io/docs/"><img align="right" src="https://img.shields.io/badge/documentation-gray"></a>
<br>
<p align="center">
  <img src="https://avatars1.githubusercontent.com/u/57274804?s=400&u=058242df13e206950c08efd68a540445ce4da17f&v=4" width="100">
</p>

# Confetti Framework

## About Confetti

Confetti is a lightweight web application framework that can be used in both monolithic and microservice environments. It has no external dependencies, making it perfect for microservices. On the other hand, it includes the core features you expect from a framework, so if you're familiar with monolithic frameworks like Laravel, you'll feel right at home. However, Confetti follows idiomatic Go principles and encourages best practices in Go development.

## Get Started

Welcome to your Confetti application! This repository serves as a template for your new project, providing essential configuration files and a robust foundation for your codebase.

### Steps to Begin

1. **Create Your Repository:**  
   Click **"Use this template"** followed by **"Create a new repository"** to set up your project.

2. **Clone the Repository:**  
   Once your repository is created, clone it to your local machine.

3. **Launch the API Server:**  
   Navigate to your project directory and run the following command:
   ```sh
   go run cmd/api/main.go api:serve
   ```
   This command starts the API server and activates your routes.

4. **Access Your First Endpoint:**  
   Your initial endpoint is now live! Visit [http://localhost:8080/status](http://localhost:8080/status) to see it in action.

### File Structure

- **cmd**: Contains application commands.
- **internal**: Includes files used only within the directory.
- **external**: Can be used throughout the entire application.
- **config**: Stores configuration files.
- **test**: Contains feature and HTTP tests. Unit tests are located next to the tested files.

## Features

Confetti provides multiple features to serve as the foundation of your application:

- [App Server with Web Routing](#app-server-with-web-routing)
- [Middlewares](#middlewares)
- [Error Handling](#error-handling)
- [Custom Commands](#custom-commands)
- [Configuration](#configuration)

### App Server with Web Routing

You can define routes in `cmd/api/command/api_serve_command.go`:

```go
var ApiRoutes = []handler.Route{
	handler.New("GET /ping", ping.Index),
	handler.New("GET /status", status.Index).AppendMiddleware(middleware.AuthMiddleware{Permission: "Show status"}),
}
```

#### Example Controller:

```go
package status

import (
	"net/http"
	"src/internal/pkg/handler"
)

// Index returns the status of the application
func Index(response http.ResponseWriter, req *http.Request) error {
	data := make(map[string]any)
	data["status"] = "active"

	return handler.ToJson(response, data, http.StatusOK)
}
```

### Middlewares

You can register middlewares like this:

```go
handler.New("GET /status", status.Index).AppendMiddleware(middleware.AuthMiddleware{Permission: "Show status"})
```

In this example, `middleware.AuthMiddleware` with the `Permission` parameter is just an example. You can modify it according to your needs.

#### Middleware Example:

```go
package middleware

import (
	"net/http"
	"src/internal/pkg/handler"
)

type AuthMiddleware struct {
	Permission string
}

func (a AuthMiddleware) Handle(next handler.Controller) handler.Controller {
	return func(w http.ResponseWriter, req *http.Request) error {
		// Perform actions before the request, e.g., checking user permissions
		err := next(w, req)
		// Perform actions after the request, e.g., logging the response status
		return err
	}
}
```

### Error Handling

In `internal/pkg/handler/error.go`, you can define a uniform error response.

For system errors:

```go
err := handler.NewSystemError(err, "psp_connection")
```

This returns only a reference to the error report instead of exposing detailed error information.

For user errors:

```go
err := handler.NewUserError("Field name is required", 422)
```

This provides a structured response containing the error message.

You can modify `handler.apiErrorHandler` to, for example, send system errors to an external error-tracking provider instead of `stderr`.

### Custom Commands

You can create your own commands. Two commands already exist. The easiest way is to copy `cmd/api/command/route_list_command.go` and modify it.

Register your command in `cmd/api/main.go`:

```go
var commands = []Command{
	command.ApiRouteList{},
	command.AppServe{},
	// Add your custom commands here
}
```

#### Command Example:

```go
package command

import (
	"fmt"
	"src/internal/pkg/handler"
)

type ApiRouteList struct {}

func (s ApiRouteList) Name() string {
	return "api:route_list"
}

func (s ApiRouteList) Description() string {
	return "Show a list of all API HTTP endpoints"
}

func (s ApiRouteList) Handle() error {
	fmt.Printf("\u001B[32mAll API endpoints:\u001B[0m\n")

	endpoints := handler.AppendApiByPath(ApiRoutes)
	for _, endpoint := range endpoints {
		fmt.Printf("%s\n", endpoint.Pattern)
	}

	return nil
}
```

### Configuration

```go
package config

var Features = struct {
	ShowHeader bool
}{
	ShowHeader: true,
}
```

Set the ShowHeader config by environment
  
```go
package config

var Features = struct {
	ShowHeader bool
}{
	ShowHeader: os.Getenv("SHOW_HEADER") == "true",
}
```
  
You can use it like this:

```go
if config.Features.ShowTest {
	// Your logic here
}
```

## Contributing

We welcome contributions! However, to keep Confetti simple and lightweight, small features should be minimal in complexity. Large feature additions are welcome but may not be merged into the core framework. Instead, we will link to your pull request or external repository in the [Features](#features) section.

### How to Contribute

1. Fork the repository.
2. Create a new branch: `git checkout -b feature-name`.
3. Implement your feature or fix.
4. Write tests if necessary.
5. Run tests to ensure everything works: `go test ./...`
6. Submit a pull request with a clear description.

Thank you for helping improve Confetti!

## License

Confetti Framework is open-source software licensed under the [MIT License](https://opensource.org/licenses/MIT).
