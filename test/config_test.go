package test

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"lanvard/app/http/decorator"
	"lanvard/bootstrap"
	"lanvard/config"
	"os"
	"reflect"
	"testing"
)

func Test_get_config(t *testing.T) {
	_ = bootstrap.NewAppFromBootApp()

	assert.IsType(t, reflect.String, support.Type(config.App.Name))
}

func Test_get_environments_for_testing(t *testing.T) {
	_ = decorator.Bootstrap(bootstrap.NewAppFromBootApp())

	assert.Equal(t, "TestLanvard", os.Getenv("APP_NAME"))
	assert.Equal(t, "TestLanvard", config.App.Name)
}
