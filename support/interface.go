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

func Empty(val interface{}) bool {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

func Dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
