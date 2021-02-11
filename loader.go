package envconfig

import (
	"fmt"
	"math/bits"
	"os"
	"reflect"
	"strconv"
)

const (
	tagKey = "env"
)

// Load loads environment variables and stores the result in the value pointed by conf.
func Load(conf interface{}) error {
	rconf := reflect.ValueOf(conf)
	if rconf.Kind() != reflect.Ptr || rconf.IsNil() || rconf.Elem().Kind() != reflect.Struct {
		return &InvalidLoadError{reflect.TypeOf(conf)}
	}

	return load(rconf.Elem())
}

func load(rv reflect.Value) error {
	for i := 0; i < rv.NumField(); i++ {
		stf := rv.Type().Field(i)
		frv := rv.Field(i)

		switch frv.Kind() {
		case reflect.Ptr:
			switch frv.Type().Elem().Kind() {
			case reflect.Struct:
				preparePtr(frv)
				load(frv.Elem())

			case reflect.String:
				preparePtr(frv)
				lookupAndSetString(stf, frv.Elem())

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				preparePtr(frv)
				if err := lookupAndSetInt(stf, frv.Elem()); err != nil {
					return err
				}

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				preparePtr(frv)
				if err := lookupAndSetUint(stf, frv.Elem()); err != nil {
					return err
				}

			case reflect.Float32, reflect.Float64:
				preparePtr(frv)
				if err := lookupAndSetFloat(stf, frv.Elem()); err != nil {
					return err
				}

			default:
				continue
			}

		case reflect.Struct:
			if err := load(frv); err != nil {
				return err
			}

		case reflect.String:
			lookupAndSetString(stf, frv)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if err := lookupAndSetInt(stf, frv); err != nil {
				return err
			}

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if err := lookupAndSetUint(stf, frv); err != nil {
				return err
			}

		case reflect.Float32, reflect.Float64:
			if err := lookupAndSetFloat(stf, frv); err != nil {
				return err
			}

		default:
			continue
		}
	}

	return nil
}

func preparePtr(rv reflect.Value) {
	if rv.IsNil() {
		rv.Set(reflect.New(rv.Type().Elem()))
	}
}

func lookupAndSetString(stf reflect.StructField, rv reflect.Value) {
	_, s, ok := lookup(stf)
	if !ok {
		return
	}

	rv.SetString(s)
}

func lookupAndSetInt(stf reflect.StructField, rv reflect.Value) error {
	k, s, ok := lookup(stf)
	if !ok {
		return nil
	}

	i, err := strconv.ParseInt(s, 10, bits.UintSize)
	if err != nil {
		return wrapParseError(k, rv, err)
	}

	rv.SetInt(i)

	return nil
}

func lookupAndSetUint(stf reflect.StructField, rv reflect.Value) error {
	k, s, ok := lookup(stf)
	if !ok {
		return nil
	}

	i, err := strconv.ParseUint(s, 10, bits.UintSize)
	if err != nil {
		return wrapParseError(k, rv, err)
	}

	rv.SetUint(i)

	return nil
}

func lookupAndSetFloat(stf reflect.StructField, rv reflect.Value) error {
	k, s, ok := lookup(stf)
	if !ok {
		return nil
	}

	f, err := strconv.ParseFloat(s, bits.UintSize)
	if err != nil {
		return wrapParseError(k, rv, err)
	}

	rv.SetFloat(f)

	return nil
}

func lookup(stf reflect.StructField) (string, string, bool) {
	k, ok := stf.Tag.Lookup(tagKey)
	if !ok {
		return "", "", false
	}

	s, ok := os.LookupEnv(k)
	if !ok {
		return k, "", false
	}

	return k, s, true
}

func wrapParseError(k string, rv reflect.Value, err error) error {
	return fmt.Errorf("envconfig: failed to parse %s as %s: %w", k, rv.Kind().String(), err)
}
