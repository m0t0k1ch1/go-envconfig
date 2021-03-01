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

var (
	underInt8String = "-129"
	minInt8String   = strconv.Itoa(math.MinInt8)
	maxInt8String   = strconv.Itoa(math.MaxInt8)
	overInt8String  = "128"

	underInt16String = "-32769"
	minInt16String   = strconv.Itoa(math.MinInt16)
	maxInt16String   = strconv.Itoa(math.MaxInt16)
	overInt16String  = "32768"

	underInt32String = "-2147483649"
	minInt32String   = strconv.Itoa(math.MinInt32)
	maxInt32String   = strconv.Itoa(math.MaxInt32)
	overInt32String  = "2147483648"

	underInt64String = "-9223372036854775809"
	minInt64String   = strconv.Itoa(math.MinInt64)
	maxInt64String   = strconv.Itoa(math.MaxInt64)
	overInt64String  = "9223372036854775808"

	maxUint8String  = strconv.FormatUint(uint64(math.MaxUint8), 10)
	overUint8String = "256"

	maxUint16String  = strconv.FormatUint(uint64(math.MaxUint16), 10)
	overUint16String = "65536"

	maxUint32String  = strconv.FormatUint(uint64(math.MaxUint32), 10)
	overUint32String = "4294967296"

	maxUint64String  = strconv.FormatUint(uint64(math.MaxUint64), 10)
	overUint64String = "18446744073709551616"
)

func TestParseFailedWithInvalidArgError(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var iaerr *InvalidArgError
		err := Parse(testEnvKey, nil)
		testutils.Equal(t, true, errors.As(err, &iaerr))
		testutils.Contains(t, err.Error(), "v cannot be nil")
	})

	t.Run("non-pointer string", func(t *testing.T) {
		var iaerr *InvalidArgError
		err := Parse(testEnvKey, string(""))
		testutils.Equal(t, true, errors.As(err, &iaerr))
		testutils.Contains(t, err.Error(), "v cannot be non-pointer string")
	})

	t.Run("nil *string", func(t *testing.T) {
		var iaerr *InvalidArgError
		err := Parse(testEnvKey, (*string)(nil))
		testutils.Equal(t, true, errors.As(err, &iaerr))
		testutils.Contains(t, err.Error(), "v cannot be nil *string")
	})
}

func TestParseFailedWithUnsupportedTypeError(t *testing.T) {
	var f float64
	var uterr *UnsupportedTypeError
	err := Parse(testEnvKey, &f)
	testutils.Equal(t, true, errors.As(err, &uterr))
	testutils.Contains(t, err.Error(), "unsupported type: float64")
}

func TestParseFailedWithNotPresentError(t *testing.T) {
	var s string
	var nperr *NotPresentError
	err := Parse(testEnvKey, &s)
	testutils.Equal(t, true, errors.As(err, &nperr))
	testutils.Contains(t, err.Error(), fmt.Sprintf("%s is not present", testEnvKey))
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
		in:  minIntString(),
		out: minInt(),
	}, {
		in:  maxIntString(),
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

func TestParseAsIntFailedWithParseError(t *testing.T) {
	cases := []struct {
		in  string
		err string
	}{{
		in:  "zero",
		err: "invalid syntax",
	}, {
		in:  underIntString(),
		err: "value out of range",
	}, {
		in:  overIntString(),
		err: "value out of range",
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var i int
			var perr *ParseError
			var nerr *strconv.NumError
			err := Parse(testEnvKey, &i)
			testutils.Equal(t, true, errors.As(err, &perr))
			testutils.Equal(t, true, errors.As(err, &nerr))
			testutils.Contains(t, err.Error(), c.err)
		})
	}
}

func TestParseAsInt8(t *testing.T) {
	cases := []struct {
		in  string
		out int8
	}{{
		in:  minInt8String,
		out: math.MinInt8,
	}, {
		in:  maxInt8String,
		out: math.MaxInt8,
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var i int8
			if err := Parse(testEnvKey, &i); err != nil {
				t.Error(err)
			} else {
				testutils.Equal(t, c.out, i)
			}
		})
	}
}

func TestParseAsInt8FailedWithParseError(t *testing.T) {
	cases := []struct {
		in  string
		err string
	}{{
		in:  "zero",
		err: "invalid syntax",
	}, {
		in:  underInt8String,
		err: "value out of range",
	}, {
		in:  overInt8String,
		err: "value out of range",
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var i int8
			var perr *ParseError
			var nerr *strconv.NumError
			err := Parse(testEnvKey, &i)
			testutils.Equal(t, true, errors.As(err, &perr))
			testutils.Equal(t, true, errors.As(err, &nerr))
			testutils.Contains(t, err.Error(), c.err)
		})
	}
}

func TestParseAsInt16(t *testing.T) {
	cases := []struct {
		in  string
		out int16
	}{{
		in:  minInt16String,
		out: math.MinInt16,
	}, {
		in:  maxInt16String,
		out: math.MaxInt16,
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var i int16
			if err := Parse(testEnvKey, &i); err != nil {
				t.Error(err)
			} else {
				testutils.Equal(t, c.out, i)
			}
		})
	}
}

func TestParseAsInt16FailedWithParseError(t *testing.T) {
	cases := []struct {
		in  string
		err string
	}{{
		in:  "zero",
		err: "invalid syntax",
	}, {
		in:  underInt16String,
		err: "value out of range",
	}, {
		in:  overInt16String,
		err: "value out of range",
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var i int16
			var perr *ParseError
			var nerr *strconv.NumError
			err := Parse(testEnvKey, &i)
			testutils.Equal(t, true, errors.As(err, &perr))
			testutils.Equal(t, true, errors.As(err, &nerr))
			testutils.Contains(t, err.Error(), c.err)
		})
	}
}

func TestParseAsInt32(t *testing.T) {
	cases := []struct {
		in  string
		out int32
	}{{
		in:  minInt32String,
		out: math.MinInt32,
	}, {
		in:  maxInt32String,
		out: math.MaxInt32,
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var i int32
			if err := Parse(testEnvKey, &i); err != nil {
				t.Error(err)
			} else {
				testutils.Equal(t, c.out, i)
			}
		})
	}
}

func TestParseAsInt32FailedWithParseError(t *testing.T) {
	cases := []struct {
		in  string
		err string
	}{{
		in:  "zero",
		err: "invalid syntax",
	}, {
		in:  underInt32String,
		err: "value out of range",
	}, {
		in:  overInt32String,
		err: "value out of range",
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var i int32
			var perr *ParseError
			var nerr *strconv.NumError
			err := Parse(testEnvKey, &i)
			testutils.Equal(t, true, errors.As(err, &perr))
			testutils.Equal(t, true, errors.As(err, &nerr))
			testutils.Contains(t, err.Error(), c.err)
		})
	}
}

func TestParseAsInt64(t *testing.T) {
	cases := []struct {
		in  string
		out int64
	}{{
		in:  minInt64String,
		out: math.MinInt64,
	}, {
		in:  maxInt64String,
		out: math.MaxInt64,
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var i int64
			if err := Parse(testEnvKey, &i); err != nil {
				t.Error(err)
			} else {
				testutils.Equal(t, c.out, i)
			}
		})
	}
}

func TestParseAsInt64FailedWithParseError(t *testing.T) {
	cases := []struct {
		in  string
		err string
	}{{
		in:  "zero",
		err: "invalid syntax",
	}, {
		in:  underInt64String,
		err: "value out of range",
	}, {
		in:  overInt64String,
		err: "value out of range",
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var i int64
			var perr *ParseError
			var nerr *strconv.NumError
			err := Parse(testEnvKey, &i)
			testutils.Equal(t, true, errors.As(err, &perr))
			testutils.Equal(t, true, errors.As(err, &nerr))
			testutils.Contains(t, err.Error(), c.err)
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
		in:  maxUintString(),
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

func TestParseAsUintFailedWithParseError(t *testing.T) {
	cases := []struct {
		in  string
		err string
	}{{
		in:  "zero",
		err: "invalid syntax",
	}, {
		in:  overUintString(),
		err: "value out of range",
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var u uint
			var perr *ParseError
			var nerr *strconv.NumError
			err := Parse(testEnvKey, &u)
			testutils.Equal(t, true, errors.As(err, &perr))
			testutils.Equal(t, true, errors.As(err, &nerr))
			testutils.Contains(t, err.Error(), c.err)
		})
	}
}

func TestParseAsUint8(t *testing.T) {
	cases := []struct {
		in  string
		out uint8
	}{{
		in:  "0",
		out: 0,
	}, {
		in:  maxUint8String,
		out: math.MaxUint8,
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var u uint8
			if err := Parse(testEnvKey, &u); err != nil {
				t.Error(err)
			} else {
				testutils.Equal(t, c.out, u)
			}
		})
	}
}

func TestParseAsUint8FailedWithParseError(t *testing.T) {
	cases := []struct {
		in  string
		err string
	}{{
		in:  "zero",
		err: "invalid syntax",
	}, {
		in:  overUint8String,
		err: "value out of range",
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var u uint8
			var perr *ParseError
			var nerr *strconv.NumError
			err := Parse(testEnvKey, &u)
			testutils.Equal(t, true, errors.As(err, &perr))
			testutils.Equal(t, true, errors.As(err, &nerr))
			testutils.Contains(t, err.Error(), c.err)
		})
	}
}

func TestParseAsUint16(t *testing.T) {
	cases := []struct {
		in  string
		out uint16
	}{{
		in:  "0",
		out: 0,
	}, {
		in:  maxUint16String,
		out: math.MaxUint16,
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var u uint16
			if err := Parse(testEnvKey, &u); err != nil {
				t.Error(err)
			} else {
				testutils.Equal(t, c.out, u)
			}
		})
	}
}

func TestParseAsUint16FailedWithParseError(t *testing.T) {
	cases := []struct {
		in  string
		err string
	}{{
		in:  "zero",
		err: "invalid syntax",
	}, {
		in:  overUint16String,
		err: "value out of range",
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var u uint16
			var perr *ParseError
			var nerr *strconv.NumError
			err := Parse(testEnvKey, &u)
			testutils.Equal(t, true, errors.As(err, &perr))
			testutils.Equal(t, true, errors.As(err, &nerr))
			testutils.Contains(t, err.Error(), c.err)
		})
	}
}

func TestParseAsUint32(t *testing.T) {
	cases := []struct {
		in  string
		out uint32
	}{{
		in:  "0",
		out: 0,
	}, {
		in:  maxUint32String,
		out: math.MaxUint32,
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var u uint32
			if err := Parse(testEnvKey, &u); err != nil {
				t.Error(err)
			} else {
				testutils.Equal(t, c.out, u)
			}
		})
	}
}

func TestParseAsUint32FailedWithParseError(t *testing.T) {
	cases := []struct {
		in  string
		err string
	}{{
		in:  "zero",
		err: "invalid syntax",
	}, {
		in:  overUint32String,
		err: "value out of range",
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var u uint32
			var perr *ParseError
			var nerr *strconv.NumError
			err := Parse(testEnvKey, &u)
			testutils.Equal(t, true, errors.As(err, &perr))
			testutils.Equal(t, true, errors.As(err, &nerr))
			testutils.Contains(t, err.Error(), c.err)
		})
	}
}

func TestParseAsUint64(t *testing.T) {
	cases := []struct {
		in  string
		out uint64
	}{{
		in:  "0",
		out: 0,
	}, {
		in:  maxUint64String,
		out: math.MaxUint64,
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var u uint64
			if err := Parse(testEnvKey, &u); err != nil {
				t.Error(err)
			} else {
				testutils.Equal(t, c.out, u)
			}
		})
	}
}

func TestParseAsUint64FailedWithParseError(t *testing.T) {
	cases := []struct {
		in  string
		err string
	}{{
		in:  "zero",
		err: "invalid syntax",
	}, {
		in:  overUint64String,
		err: "value out of range",
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var u uint64
			var perr *ParseError
			var nerr *strconv.NumError
			err := Parse(testEnvKey, &u)
			testutils.Equal(t, true, errors.As(err, &perr))
			testutils.Equal(t, true, errors.As(err, &nerr))
			testutils.Contains(t, err.Error(), c.err)
		})
	}
}

func TestParseAsBool(t *testing.T) {
	cases := []struct {
		in  string
		out bool
	}{{
		in:  "",
		out: false,
	}, {
		in:  "0",
		out: false,
	}, {
		in:  "1",
		out: true,
	}, {
		in:  "false",
		out: true,
	}, {
		in:  "true",
		out: true,
	}}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			os.Setenv(testEnvKey, c.in)
			defer os.Clearenv()

			var b bool
			if err := Parse(testEnvKey, &b); err != nil {
				t.Error(err)
			} else {
				testutils.Equal(t, c.out, b)
			}
		})
	}
}

func underIntString() string {
	if bits.UintSize == 32 {
		return underInt32String
	}

	return underInt64String
}

func minInt() int {
	if bits.UintSize == 32 {
		return math.MinInt32
	}

	return math.MinInt64
}

func minIntString() string {
	return strconv.Itoa(minInt())
}

func maxInt() int {
	if bits.UintSize == 32 {
		return math.MaxInt32
	}

	return math.MaxInt64
}

func maxIntString() string {
	return strconv.Itoa(maxInt())
}

func overIntString() string {
	if bits.UintSize == 32 {
		return overInt32String
	}

	return overInt64String
}

func maxUint() uint {
	if bits.UintSize == 32 {
		return math.MaxUint32
	}

	return math.MaxUint64
}

func maxUintString() string {
	return strconv.FormatUint(uint64(maxUint()), 10)
}

func overUintString() string {
	if bits.UintSize == 32 {
		return overUint32String
	}

	return overUint64String
}
