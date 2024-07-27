package config

import "time"

var AppServe = struct {
	Ssl     bool
	Host    string
	Port    int
	Timeout time.Duration
}{
	Ssl:     EnvBoolOr("APP_SSL", false),
	Host:    EnvStringOr("APP_HOST", ""),
	Port:    EnvIntOr("APP_PORT", 8080),
	Timeout: time.Duration(EnvIntOr("APP_TIMEOUT", 30)) * time.Second,
}
