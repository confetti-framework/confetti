package test

import (
	"github.com/stretchr/testify/assert"
	"lanvard/src/app/http"
	"lanvard/src/bootstrap"
	"os"
	"testing"
)

func Test_bootstrap_environments(t *testing.T) {

	_ = http.Bootstrap(bootstrap.App())

	assert.NotZero(t, os.Getenv("APP_ENV"))
}
