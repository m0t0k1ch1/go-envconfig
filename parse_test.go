package envparser

import (
	"os"
	"testing"

	"github.com/m0t0k1ch1/go-envparser/internal/testutils"
)

const (
	testEnvKey = "GO_ENVPARSER_TEST"
)

func TestParseAsString(t *testing.T) {
	defer os.Clearenv()
	os.Setenv(testEnvKey, "string")

	var s string
	if err := Parse(testEnvKey, &s); err != nil {
		t.Error(err)
		return
	}

	testutils.Equal(t, s, "string")
}
