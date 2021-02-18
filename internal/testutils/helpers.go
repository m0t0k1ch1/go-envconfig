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

// Contains checks if s contains substr.
func Contains(t *testing.T, s string, substr string) {
	t.Helper()

	if ok := strings.Contains(s, substr); !ok {
		t.Errorf(`"%s" does not contain "%s"`, s, substr)
	}
}
