package testutils

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Equal checks if want and got are equal.
func Equal(t *testing.T, want, got interface{}) {
	t.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("diff: %s", diff)
	}
}

// TestErrorMessage checks if err message contains s.
func TestErrorMessage(t *testing.T, s string, err error) {
	t.Helper()

	if s == "" {
		if err != nil {
			t.Errorf("err is not nil: %v", err)
			return
		}
		return
	}

	if contained := strings.Contains(err.Error(), s); !contained {
		t.Errorf(`err message does not contain "%s": %v`, s, err)
	}
}
