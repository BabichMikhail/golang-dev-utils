package utils

import (
	"reflect"
)

func GetTypeName(value interface{}) string {
	t := reflect.TypeOf(value)
	return Ternary(t.Kind() == reflect.Ptr, "*" + t.Elem().Name(), t.Name()).(string)
}
