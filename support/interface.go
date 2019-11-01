package support

import (
	"reflect"
)

func Name(element interface{}) string  {
	if Type(element) == reflect.String {
		return element.(string)
	}

	if Type(element) == reflect.Struct {
		return reflect.TypeOf(element).String()
	}

	return reflect.TypeOf(element).Elem().String()
}

func Type(element interface{}) interface{} {

	if element == nil {
		return reflect.TypeOf(&element).Kind()
	}

	return reflect.TypeOf(element).Kind()
}