package utils

import (
	"fmt"
	"reflect"
)

// Caution: Make sure the obj parameter is a struct value and not a pointer to a struct
func PrintStructValues(obj interface{}) {
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	fmt.Println("Fields of", typ.Name()+":")
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		fieldValue := field.Interface()

		fmt.Println(fieldName+":", fieldValue)
	}
}
