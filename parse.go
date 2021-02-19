package envparser

import (
	"math/bits"
	"os"
	"reflect"
	"strconv"
)

var (
	supportedKinds = []reflect.Kind{
		reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
	}
)

// Parse parses the environment variable named by k and stores the result in the value pointed by v.
func Parse(k string, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidArgError{reflect.TypeOf(v)}
	}
	if !isSupportedKind(rv.Elem().Kind()) {
		return &UnsupportedTypeError{rv.Elem().Type()}
	}

	p := &parser{k, rv.Elem(), nil}

	return p.parse()
}

func isSupportedKind(kind reflect.Kind) bool {
	for _, supportedKind := range supportedKinds {
		if supportedKind == kind {
			return true
		}
	}

	return false
}

type parser struct {
	k   string
	rv  reflect.Value
	err error
}

func (p *parser) parse() error {
	s, ok := os.LookupEnv(p.k)
	if !ok {
		return &NotPresentError{p.k}
	}

	switch p.rv.Kind() {
	case reflect.String:
		p.rv.SetString(s)
	case reflect.Int:
		p.parseAndSetInt(s, bits.UintSize)
	case reflect.Int8:
		p.parseAndSetInt(s, 8)
	case reflect.Int16:
		p.parseAndSetInt(s, 16)
	case reflect.Int32:
		p.parseAndSetInt(s, 32)
	case reflect.Int64:
		p.parseAndSetInt(s, 64)
	case reflect.Uint:
		p.parseAndSetUint(s, bits.UintSize)
	case reflect.Uint8:
		p.parseAndSetUint(s, 8)
	case reflect.Uint16:
		p.parseAndSetUint(s, 16)
	case reflect.Uint32:
		p.parseAndSetUint(s, 32)
	case reflect.Uint64:
		p.parseAndSetUint(s, 64)
	default:
		p.err = &UnsupportedTypeError{p.rv.Type()}
	}

	return p.err
}

func (p *parser) parseAndSetInt(s string, bitSize int) {
	i, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		p.err = &ParseError{err, p.k, p.rv.Type()}
	}
	p.rv.SetInt(i)
}

func (p *parser) parseAndSetUint(s string, bitSize int) {
	u, err := strconv.ParseUint(s, 10, bitSize)
	if err != nil {
		p.err = &ParseError{err, p.k, p.rv.Type()}
	}
	p.rv.SetUint(u)
}
