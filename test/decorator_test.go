package test

import (
	"github.com/stretchr/testify/assert"
	"lanvard/app/http/decorator"
	"lanvard/bootstrap"
	"lanvard/config"
	"os"
	"testing"
)

func Test_bootstrap_environments(t *testing.T) {
	_ = decorator.Bootstrap(bootstrap.NewApp())

	assert.NotZero(t, os.Getenv("APP_ENV"))
}

func Test_bootstrap_environments_for_testing(t *testing.T) {
	_ = decorator.Bootstrap(bootstrap.NewApp())

	assert.Equal(t, "TestLanvard", os.Getenv("APP_NAME"))
	assert.Equal(t, "TestLanvard", config.App.Name)
}
