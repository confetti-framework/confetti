package test

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"lanvard/bootstrap"
	"lanvard/config"
	"reflect"
	"testing"
)

func Test_get_config(t *testing.T) {
	_ = bootstrap.NewApp()

	assert.IsType(t, reflect.String, support.Type(config.App.Name))
}
