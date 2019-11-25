package support

import (
	"fmt"
	"reflect"
)

func Name(element interface{}) string {
	if Type(element) == reflect.String {
		return element.(string)
	}

	if Type(element) == reflect.Struct {
		return reflect.TypeOf(element).String()
	}

	return reflect.TypeOf(element).Elem().String()
}

func Package(element interface{}) string {

	if element == nil {
		return reflect.TypeOf(&element).Elem().PkgPath()
	}

	return reflect.TypeOf(element).Elem().PkgPath()
}

func Type(element interface{}) interface{} {

	if element == nil {
		return reflect.TypeOf(&element).Kind()
	}

	return reflect.TypeOf(element).Kind()
}

func Dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
