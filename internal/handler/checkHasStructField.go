package handler

import (
	"reflect"
)

func HasField(s interface{}, fieldName string) bool {
	// Get the reflect.Value of the interface
	v := reflect.ValueOf(s)

	// Ensure it's a struct (or a pointer to a struct)
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // Dereference the pointer
	}
	if v.Kind() != reflect.Struct {
		return false // Not a struct
	}

	// Look for the field by name
	field := v.FieldByName(fieldName)
	return field.IsValid() // IsValid returns true if the field exists
}
