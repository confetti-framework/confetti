package test

import (
	"github.com/stretchr/testify/assert"
	kernelAlias "lanvard/contract/http"
	"lanvard/support"
	"reflect"
	"testing"
)

type testInterface interface{}
type testStruct struct{}

func TestNameFromNil(t *testing.T) {
	name := support.Name((*testInterface)(nil))
	assert.Equal(t, "test.testInterface", name)
}

func TestNameFromStruct(t *testing.T) {
	name := support.Name(testStruct{})
	assert.Equal(t, "test.testStruct", name)
}

func TestNameFromString(t *testing.T) {
	name := support.Name("InterfaceByString")
	assert.Equal(t, "InterfaceByString", name)
}

func TestNameWithAlias(t *testing.T) {
	name := support.Name((*kernelAlias.Kernel)(nil))
	assert.Equal(t, "http.Kernel", name)
}

func TestTypeFromInterface(t *testing.T) {
	reflectType := support.Type((*testInterface)(nil))
	assert.Equal(t, reflect.Ptr, reflectType)
}

func TestTypeFromString(t *testing.T) {
	reflectType := support.Type("string")
	assert.Equal(t, reflect.String, reflectType)
}
