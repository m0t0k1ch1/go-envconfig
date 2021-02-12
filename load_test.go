package envconfig

import (
	"os"
	"testing"

	"github.com/m0t0k1ch1/go-envconfig/internal/testutils"
)

var testConfSample = testConfig{
	String: "string",
	StringPtr: func() *string {
		s := "string_ptr"
		return &s
	}(),
	Int: 1001,
	IntPtr: func() *int {
		i := int(1002)
		return &i
	}(),
	Uint: 2001,
	UintPtr: func() *uint {
		i := uint(2002)
		return &i
	}(),
	Float32: 32.1,
	Float32Ptr: func() *float32 {
		f := float32(32.2)
		return &f
	}(),
	Float64: 64.1,
	Float64Ptr: func() *float64 {
		f := float64(64.2)
		return &f
	}(),
	Struct: testStruct1{
		String: "struct1_string",
	},
	StructPtr: &testStruct2{
		String: "struct2_string",
	},
}

type testConfig struct {
	String     string   `env:"STRING"`
	StringPtr  *string  `env:"STRING_PTR"`
	Int        int      `env:"INT"`
	IntPtr     *int     `env:"INT_PTR"`
	Uint       uint     `env:"UINT"`
	UintPtr    *uint    `env:"UINT_PTR"`
	Float32    float32  `env:"FLOAT32"`
	Float32Ptr *float32 `env:"FLOAT32_PTR"`
	Float64    float64  `env:"FLOAT64"`
	Float64Ptr *float64 `env:"FLOAT64_PTR"`
	Struct     testStruct1
	StructPtr  *testStruct2
}

type testStruct1 struct {
	String string `env:"STRUCT1_STRING"`
}

type testStruct2 struct {
	String string `env:"STRUCT2_STRING"`
}

func TestLoad(t *testing.T) {
	defer os.Clearenv()

	os.Setenv("STRING", "string")
	os.Setenv("STRING_PTR", "string_ptr")
	os.Setenv("INT", "1001")
	os.Setenv("INT_PTR", "1002")
	os.Setenv("UINT", "2001")
	os.Setenv("UINT_PTR", "2002")
	os.Setenv("FLOAT32", "32.1")
	os.Setenv("FLOAT32_PTR", "32.2")
	os.Setenv("FLOAT64", "64.1")
	os.Setenv("FLOAT64_PTR", "64.2")
	os.Setenv("STRUCT1_STRING", "struct1_string")
	os.Setenv("STRUCT2_STRING", "struct2_string")

	var conf testConfig
	if err := Load(&conf); err != nil {
		t.Error(err)
		return
	}

	testutils.Equal(t, conf, testConfSample)
}
