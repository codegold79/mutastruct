/*
Package mutastruct uses reflect to will look if a field exists in a struct and
will set it if the type is right.
*/
package mutastruct

import (
	"fmt"
	"reflect"
)

// Set will set a field value if the field exists and the value is the correct type.
// The first argument to Set is a pointer to the struct that is to be modified.
func Set(anyStruct any, fieldName string, setting any) error {
	anyStructValue := reflect.ValueOf(anyStruct)

	// Verify that the first argument is a pointer to a struct.
	if anyStructValue.Kind() != reflect.Pointer {
		return fmt.Errorf("expected pointer to a struct: %w", ErrTypeWrongType)
	}
	if anyStructValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected struct: %w", ErrTypeWrongType)
	}

	// Verify field exists in struct.
	foundField := anyStructValue.Elem().FieldByName(fieldName)
	if (foundField == reflect.Value{}) {
		return fmt.Errorf("the passed in struct (%v) does not have field %q: %w", anyStruct, fieldName, ErrTypeStructFieldNotFound)
	}

	// Verify field can be changed.
	if !foundField.CanSet() {
		return fmt.Errorf("the field %s to be changed must be exported (public): %w", fieldName, ErrTypeFieldNotSettable)
	}

	// Verify the desired fieldValue is the same type as the given struct's
	// field type.
	settingValue := reflect.ValueOf(setting)
	if foundField.Kind() != settingValue.Kind() {
		return fmt.Errorf("the desired field value (%s) does not match the struct's field type (%s): %w", settingValue.Kind(), foundField.Kind(), ErrTypeStructFieldIncompatible)
	}

	// Set the field to the desired value.
	reflect.ValueOf(anyStruct).Elem().FieldByName(fieldName).Set(settingValue)
	return nil
}
