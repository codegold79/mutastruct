package mutastruct

import (
	"errors"
	"testing"
)

func TestSetStringField(t *testing.T) {
	type TestStruct struct {
		Str string
	}

	var testStruct TestStruct

	desc := "Set is able to a string in an existing string-type field"
	strVal := "string for testing"
	want := TestStruct{strVal}

	t.Run(desc, func(t *testing.T) {
		if err := Set(&testStruct, "Str", strVal); err != nil {
			t.Error(err)
		}

		if testStruct != want {
			t.Errorf("got: %v of type: %T, want: %v of type: %T", testStruct, testStruct, want, want)
		}
	})

	desc = "Set requires a pointer to a struct in which to set a field"
	t.Run(desc, func(t *testing.T) {
		if err := Set(testStruct, "Str", strVal); !errors.Is(err, ErrTypeWrongType) {
			t.Errorf("got %v, want error: %v", err, ErrTypeWrongType)
		}

		s := "first arg is not a struct"
		if err := Set(&s, "no field", strVal); !errors.Is(err, ErrTypeWrongType) {
			t.Errorf("got %v, want error: %v", err, ErrTypeWrongType)
		}
	})

	desc = "Set is unable find field to set"
	t.Run(desc, func(t *testing.T) {
		if err := Set(&testStruct, "invisible", strVal); !errors.Is(err, ErrTypeStructFieldNotFound) {
			t.Errorf("got: %v, want error: %v", err, ErrTypeStructFieldNotFound)
		}
	})

	desc = "Set is unable to set string field to bool value"
	boolVal := false
	t.Run(desc, func(t *testing.T) {
		if err := Set(&testStruct, "Str", boolVal); !errors.Is(err, ErrTypeStructFieldIncompatible) {
			t.Errorf("got: %v, want error: %v", err, ErrTypeStructFieldIncompatible)
		}
	})

	desc = "Set is unable to set string field to int value"
	intVal := 10000
	t.Run(desc, func(t *testing.T) {
		if err := Set(&testStruct, "Str", intVal); !errors.Is(err, ErrTypeStructFieldIncompatible) {
			t.Errorf("got: %v, want error: %v", err, ErrTypeStructFieldIncompatible)
		}
	})

	desc = "Set is unable to set string field to struct value"
	structVal := TestStruct{}
	t.Run(desc, func(t *testing.T) {
		if err := Set(&testStruct, "Str", structVal); !errors.Is(err, ErrTypeStructFieldIncompatible) {
			t.Errorf("got: %v, want error: %v", err, ErrTypeStructFieldIncompatible)
		}
	})
}

func TestSetBoolField(t *testing.T) {
	type TestStruct struct {
		Bool bool
	}

	var testStruct TestStruct

	desc := "Set is able to a string in an existing string-type field"
	boolVal := true
	want := TestStruct{boolVal}

	t.Run(desc, func(t *testing.T) {
		if err := Set(&testStruct, "Bool", boolVal); err != nil {
			t.Error(err)
		}

		if testStruct != want {
			t.Errorf("got: %v of type: %T, want: %v of type: %T", testStruct, testStruct, want, want)
		}
	})
}

func TestSetIntField(t *testing.T) {
	type TestStruct struct {
		Int int
	}

	var testStruct TestStruct

	desc := "Set is able to a string in an existing string-type field"
	intVal := 100
	want := TestStruct{intVal}

	t.Run(desc, func(t *testing.T) {
		if err := Set(&testStruct, "Int", intVal); err != nil {
			t.Error(err)
		}

		if testStruct != want {
			t.Errorf("got: %v of type: %T, want: %v of type: %T", testStruct, testStruct, want, want)
		}
	})
}
