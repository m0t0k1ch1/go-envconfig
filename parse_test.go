package envparser

import (
	"errors"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strconv"
	"testing"

	"github.com/m0t0k1ch1/go-envparser/internal/testutils"
)

const (
	testEnvKey = "GO_ENVPARSER_TEST"
)

func TestParseInvalidArgError(t *testing.T) {
	var err error
	var iaerr *InvalidArgError

	err = Parse(testEnvKey, nil)
	testutils.Equal(t, true, errors.As(err, &iaerr))
	testutils.Contains(t, err.Error(), "v cannot be nil")

	err = Parse(testEnvKey, string(""))
	testutils.Equal(t, true, errors.As(err, &iaerr))
	testutils.Contains(t, err.Error(), "v cannot be non-pointer string")

	err = Parse(testEnvKey, (*string)(nil))
	testutils.Equal(t, true, errors.As(err, &iaerr))
	testutils.Contains(t, err.Error(), "v cannot be nil *string")
}

func TestParseUnsupportedTypeError(t *testing.T) {
	var err error
	var uterr *UnsupportedTypeError

	var b bool
	err = Parse(testEnvKey, &b)
	testutils.Equal(t, true, errors.As(err, &uterr))
	testutils.Contains(t, err.Error(), "unsupported type: bool")
}

func TestParseNotPresentError(t *testing.T) {
	var err error
	var nperr *NotPresentError

	var s string
	err = Parse(testEnvKey, &s)
	testutils.Equal(t, true, errors.As(err, &nperr))
	testutils.Contains(t, err.Error(), fmt.Sprintf("%s is not present", testEnvKey))
}

func TestParseError(t *testing.T) {
	var err error
	var perr *ParseError
	var nerr *strconv.NumError

	os.Setenv(testEnvKey, "zero")
	defer os.Clearenv()

	var i int
	err = Parse(testEnvKey, &i)
	testutils.Equal(t, true, errors.As(err, &perr))
	testutils.Equal(t, true, errors.As(err, &nerr))
	testutils.Contains(t, err.Error(), fmt.Sprintf("cannot parse %s as int", testEnvKey))
}

func TestParseAsString(t *testing.T) {
	in := "string"
	out := "string"

	os.Setenv(testEnvKey, in)
	defer os.Clearenv()

	var s string
	if err := Parse(testEnvKey, &s); err != nil {
		t.Error(err)
	} else {
		testutils.Equal(t, out, s)
	}
}

func TestParseAsInt(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{{
		in:  strconv.Itoa(minInt()),
		out: minInt(),
	}, {
		in:  strconv.Itoa(maxInt()),
		out: maxInt(),
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var i int
			if err := Parse(testEnvKey, &i); err != nil {
				t.Error(err)
			} else {
				testutils.Equal(t, c.out, i)
			}
		})
	}
}

func TestParseAsUint(t *testing.T) {
	cases := []struct {
		in  string
		out uint
	}{{
		in:  "0",
		out: 0,
	}, {
		in:  strconv.FormatUint(uint64(maxUint()), 10),
		out: maxUint(),
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var u uint
			if err := Parse(testEnvKey, &u); err != nil {
				t.Error(err)
			} else {
				testutils.Equal(t, c.out, u)
			}
		})
	}
}

func minInt() int {
	if bits.UintSize == 32 {
		return math.MinInt32
	}

	return math.MinInt64
}

func maxInt() int {
	if bits.UintSize == 32 {
		return math.MaxInt32
	}

	return math.MaxInt64
}

func maxUint() uint {
	if bits.UintSize == 32 {
		return math.MaxUint32
	}

	return math.MaxUint64
}
