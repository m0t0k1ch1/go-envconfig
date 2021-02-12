package testutils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Equal checks if actual and expected are equal.
func Equal(t *testing.T, actual, expected interface{}) {
	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("diff: %s", diff)
	}
}
