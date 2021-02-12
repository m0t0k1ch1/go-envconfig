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
}

type testConfig struct {
	String    string  `env:"STRING"`
	StringPtr *string `env:"STRING_PTR"`
}

func TestLoad(t *testing.T) {
	defer os.Clearenv()

	os.Setenv("STRING", "string")
	os.Setenv("STRING_PTR", "string_ptr")

	var conf testConfig
	if err := Load(&conf); err != nil {
		t.Error(err)
		return
	}

	testutils.Equal(t, conf, testConfSample)
}
