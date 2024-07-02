package config

import (
    "os"
    "strconv"
)

// EnvToString fetches the value of the environment variable named by the key.
// It panics if the variable is not found.
func EnvToString(search string) string {
    env, ok := os.LookupEnv(search)
    if !ok {
        panic("environment variable not found: " + search)
    }
    return env
}

// EnvToStringOr fetches the value of the environment variable named by the key.
// If the variable is not present, it returns the provided default value.
func EnvToStringOr(search string, def string) string {
    env, ok := os.LookupEnv(search)
    if !ok {
        return def
    }
    return env
}

// EnvToInt fetches the value of the environment variable named by the key
// and converts it to an integer. It panics if the variable is not present or the conversion fails.
func EnvToInt(search string) int {
    env, ok := os.LookupEnv(search)
    if !ok {
        panic("environment variable not found: " + search)
    }
    value, err := strconv.Atoi(env)
    if err != nil {
        panic("environment variable conversion to int failed: " + search + " with value " + env)
    }
    return value
}

// EnvToIntOr fetches the value of the environment variable named by the key
// and converts it to an integer. If the variable is not present, it returns the provided default value.
// If the conversion fails, it panics.
func EnvToIntOr(search string, def int) int {
    env, ok := os.LookupEnv(search)
    if !ok {
        return def
    }
    value, err := strconv.Atoi(env)
    if err != nil {
        panic("environment variable conversion to int failed: " + search + " with value " + env)
    }
    return value
}

// EnvToBool fetches the value of the environment variable named by the key
// and converts it to a boolean. It panics if the variable is not present or the conversion fails.
func EnvToBool(search string) bool {
    env, ok := os.LookupEnv(search)
    if !ok {
        panic("environment variable not found: " + search)
    }
    value, err := strconv.ParseBool(env)
    if err != nil {
        panic("environment variable conversion to bool failed: " + search + " with value " + env)
    }
    return value
}

// EnvToBoolOr fetches the value of the environment variable named by the key
// and converts it to a boolean. If the variable is not present, it returns the provided default value.
// If the conversion fails, it panics.
func EnvToBoolOr(search string, def bool) bool {
    env, ok := os.LookupEnv(search)
    if !ok {
        return def
    }
    value, err := strconv.ParseBool(env)
    if err != nil {
        panic("environment variable conversion to bool failed: " + search + " with value " + env)
    }
    return value
}
