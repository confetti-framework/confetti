package test

import (
	"github.com/stretchr/testify/assert"
	"lanvard/config"
	"lanvard/src/bootstrap"
	"lanvard/support"
	"reflect"
	"testing"
)

func Test_get_config(t *testing.T) {
	_ = bootstrap.App()

	assert.IsType(t, reflect.String, support.Type(config.App.Name))
}
