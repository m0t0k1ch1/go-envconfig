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
	testutils.TestErrorMessage(t, "v cannot be nil", Parse(testEnvKey, nil))
	testutils.TestErrorMessage(t, "v cannot be non-pointer string", Parse(testEnvKey, string("")))
	testutils.TestErrorMessage(t, "v cannot be nil *string", Parse(testEnvKey, (*string)(nil)))
}

func TestParseUnsupportedTypeError(t *testing.T) {
	var b bool
	testutils.TestErrorMessage(t, "unsupported type: bool", Parse(testEnvKey, &b))
}

func TestParseNotPresentError(t *testing.T) {
	var s string
	testutils.TestErrorMessage(t, fmt.Sprintf("%s is not present", testEnvKey), Parse(testEnvKey, &s))
}

func TestParseError(t *testing.T) {
	defer os.Clearenv()
	os.Setenv(testEnvKey, "zero")

	var i int
	err := Parse(testEnvKey, &i)
	testutils.TestErrorMessage(t, fmt.Sprintf("cannot parse %s as int", testEnvKey), err)

	var perr *ParseError
	testutils.Equal(t, true, errors.As(err, &perr))

	var nerr *strconv.NumError
	testutils.Equal(t, true, errors.As(err, &nerr))
}

func TestParseAsString(t *testing.T) {
	defer os.Clearenv()
	os.Setenv(testEnvKey, "string")

	var s string
	testutils.TestErrorMessage(t, "", Parse(testEnvKey, &s))
	testutils.Equal(t, "string", s)
}

func TestParseAsInt(t *testing.T) {
	cases := []struct {
		s   string
		i   int
		err string
	}{{
		s:   strconv.Itoa(minInt()),
		i:   minInt(),
		err: "",
	}, {
		s:   strconv.Itoa(maxInt()),
		i:   maxInt(),
		err: "",
	}}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			defer os.Clearenv()
			os.Setenv(testEnvKey, c.s)

			var i int
			testutils.TestErrorMessage(t, c.err, Parse(testEnvKey, &i))
			testutils.Equal(t, c.i, i)
		})
	}
}

func TestParseAsUint(t *testing.T) {
	cases := []struct {
		s   string
		u   uint
		err string
	}{{
		s:   "0",
		u:   0,
		err: "",
	}, {
		s:   strconv.FormatUint(uint64(maxUint()), 10),
		u:   maxUint(),
		err: "",
	}}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			defer os.Clearenv()
			os.Setenv(testEnvKey, c.s)

			var u uint
			testutils.TestErrorMessage(t, c.err, Parse(testEnvKey, &u))
			testutils.Equal(t, c.u, u)
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
