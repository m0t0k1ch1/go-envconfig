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
