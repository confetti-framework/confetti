package config

import "time"

var AppServe = struct {
    Ssl     bool
    Host    string
    Port    int
    Timeout time.Duration
}{
    Ssl:     EnvToBoolOr("APP_SSL", false),
    Host:    EnvToStringOr("APP_HOST", ""),
    Port:    EnvToIntOr("APP_PORT", 8080),
    Timeout: time.Duration(EnvToIntOr("APP_TIMEOUT", 30)) * time.Second,
}
