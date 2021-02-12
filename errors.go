package envconfig

import (
	"fmt"
	"reflect"
)

// InvalidLoadError describes an invalid argument passed to Load.
type InvalidLoadError struct {
	Type reflect.Type
}

func (e *InvalidLoadError) Error() string {
	if e.Type == nil {
		return "envconfig: Load(nil)"
	}
	if e.Type.Kind() != reflect.Ptr {
		return fmt.Sprintf("envconfig: Load(non-pointer %s)", e.Type.String())
	}

	return fmt.Sprintf("envconfig: Load(nil %s)", e.Type.String())
}

// InvalidTypeError describes an environment variable was not appropriate for a value of a specific Go type.
type InvalidTypeError struct {
	Key  string
	Type reflect.Type
}

func (e *InvalidTypeError) Error() string {
	return fmt.Sprintf("envconfig: cannot load %s into Go value of type %s", e.Key, e.Type.String())
}
