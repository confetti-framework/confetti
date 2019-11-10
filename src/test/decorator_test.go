package test

import (
	"github.com/stretchr/testify/assert"
	"lanvard/src/app/http/decorator"
	"lanvard/src/bootstrap"
	"os"
	"testing"
)

func Test_bootstrap_environments(t *testing.T) {

	_ = decorator.Bootstrap(bootstrap.App())

	assert.NotZero(t, os.Getenv("APP_ENV"))
}
