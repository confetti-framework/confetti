package test

import (
	"github.com/stretchr/testify/assert"
	"lanvard/app/http/decorator"
	"lanvard/bootstrap"
	"os"
	"testing"
)

func Test_bootstrap_environments(t *testing.T) {
	_ = decorator.Bootstrap(bootstrap.NewApp())

	assert.NotZero(t, os.Getenv("APP_ENV"))
}
