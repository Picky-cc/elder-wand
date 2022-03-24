package utils

import (
	"reflect"
)

func ReflectNewStruct(i interface{}) interface{} {
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		// Pointer:
		return reflect.New(reflect.ValueOf(i).Elem().Type()).Interface()
	}
	// Not pointer:
	return reflect.New(reflect.TypeOf(i)).Elem().Interface()
}
