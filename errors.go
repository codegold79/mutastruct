package mutastruct

type ErrType string

const (
	ErrTypeStructFieldNotFound     ErrType = "Field not found in struct"
	ErrTypeFieldNotSettable        ErrType = "Field not settable"
	ErrTypeWrongType               ErrType = "Wrong type"
	ErrTypeStructFieldIncompatible ErrType = "Field value is not compatible in struct"
)

func (err ErrType) Error() string {
	switch err {
	case ErrTypeWrongType:
		return "Wrong type"
	case ErrTypeStructFieldNotFound:
		return "Field not found"
	case ErrTypeStructFieldIncompatible:
		return "Field incompatible"
	case ErrTypeFieldNotSettable:
		return "Field cannot be set"
	}
	return "unknown error type"
}
