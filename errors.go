package envparser

import (
	"fmt"
	"reflect"
)

// InvalidArgError is returned when an invalid argument passed to Parse.
type InvalidArgError struct {
	Type reflect.Type
}

func (e *InvalidArgError) Error() string {
	if e.Type == nil {
		return "envparser: v can not be nil"
	}
	if e.Type.Kind() != reflect.Ptr {
		return fmt.Sprintf("envparser: v can not be non-pointer %s", e.Type.String())
	}

	return fmt.Sprintf("envparser: v can not be nil %s", e.Type.String())
}

// UnsupportedTypeError is returned when attempting to parse as an unsupported value type.
type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
	return fmt.Sprintf("envparser: unsupported type: %s", e.Type.String())
}

// NotPresentError is returned when the environment variable is not present.
type NotPresentError struct {
	Key string
}

func (e *NotPresentError) Error() string {
	return fmt.Sprintf("envparser: %s is not present", e.Key)
}

// ParseError is returned when parsing fails.
type ParseError struct {
	wrapped error
	Key     string
	Type    reflect.Type
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("envparser: cannot parse %s as %s: %v", e.Key, e.Type.String(), e.wrapped)
}

func (e *ParseError) Unwrap() error {
	return e.wrapped
}
