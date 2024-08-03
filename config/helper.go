package config

import (
	"os"
	"strconv"
)

var EnvProvider = os.LookupEnv

// EnvString fetches the value of the environment variable named by the key.
// It panics if the variable is not found.
func EnvString(search string) string {
	env, ok := EnvProvider(search)
	if !ok {
		return ""
	}
	return env
}

// EnvStringOr fetches the value of the environment variable named by the key.
// If the variable is not present, it returns the provided default value.
func EnvStringOr(search string, def string) string {
	env, ok := EnvProvider(search)
	if !ok {
		return def
	}
	return env
}

// EnvInt fetches the value of the environment variable named by the key
// and converts it to an integer. It panics if the variable is not present or the conversion fails.
func EnvInt(search string) int {
	env, ok := EnvProvider(search)
	if !ok {
		return 0
	}
	value, err := strconv.Atoi(env)
	if err != nil {
		panic("environment variable conversion to int failed: " + search + " with value " + env)
	}
	return value
}

// EnvIntOr fetches the value of the environment variable named by the key
// and converts it to an integer. If the variable is not present, it returns the provided default value.
// If the conversion fails, it panics.
func EnvIntOr(search string, def int) int {
	env, ok := EnvProvider(search)
	if !ok {
		return def
	}
	value, err := strconv.Atoi(env)
	if err != nil {
		panic("environment variable conversion to int failed: " + search + " with value " + env)
	}
	return value
}

// EnvBool fetches the value of the environment variable named by the key
// and converts it to a boolean. It panics if the variable is not present or the conversion fails.
func EnvBool(search string) bool {
	env, ok := EnvProvider(search)
	if !ok {
		return false
	}
	value, err := strconv.ParseBool(env)
	if err != nil {
		panic("environment variable conversion to bool failed: " + search + " with value " + env)
	}
	return value
}

// EnvBoolOr fetches the value of the environment variable named by the key
// and converts it to a boolean. If the variable is not present, it returns the provided default value.
// If the conversion fails, it panics.
func EnvBoolOr(search string, def bool) bool {
	env, ok := EnvProvider(search)
	if !ok {
		return def
	}
	value, err := strconv.ParseBool(env)
	if err != nil {
		panic("environment variable conversion to bool failed: " + search + " with value " + env)
	}
	return value
}
