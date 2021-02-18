package envparser

import (
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

	err = Parse(testEnvKey, nil)
	testutils.TestErrorMessage(t, "v cannot be nil", err)

	err = Parse(testEnvKey, string(""))
	testutils.TestErrorMessage(t, "v cannot be non-pointer string", err)

	err = Parse(testEnvKey, (*string)(nil))
	testutils.TestErrorMessage(t, "v cannot be nil *string", err)
}

func TestParseUnsupportedTypeError(t *testing.T) {
	var b bool
	err := Parse(testEnvKey, &b)
	testutils.TestErrorMessage(t, "unsupported type: bool", err)
}

func TestParseAsString(t *testing.T) {
	defer os.Clearenv()
	os.Setenv(testEnvKey, "string")

	var s string
	err := Parse(testEnvKey, &s)
	testutils.TestErrorMessage(t, "", err)
	testutils.Equal(t, "string", s)
}

func TestParseAsInt(t *testing.T) {
	cases := []struct {
		s   string
		i   int
		err string
	}{{
		s:   strconv.Itoa(-maxInt()),
		i:   -maxInt(),
		err: "",
	}, {
		s:   strconv.Itoa(maxInt()),
		i:   maxInt(),
		err: "",
	}, {
		s:   "zero",
		i:   0,
		err: fmt.Sprintf("cannot parse %s as int", testEnvKey),
	}}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			defer os.Clearenv()
			os.Setenv(testEnvKey, c.s)

			var i int
			err := Parse(testEnvKey, &i)
			testutils.TestErrorMessage(t, c.err, err)
			testutils.Equal(t, c.i, i)
		})
	}
}

func maxInt() int {
	if bits.UintSize == 32 {
		return math.MaxInt32
	}

	return math.MaxInt64
}
