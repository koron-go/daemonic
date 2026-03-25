// Package assert provides assert functions for testing.
package assert

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Equal[T any](tb testing.TB, want, got T, options ...cmp.Option) {
	tb.Helper()
	if d := cmp.Diff(want, got, options...); d != "" {
		tb.Errorf("assert failed, mismatch: -want +got\n%s", d)
	}
}

func IsNotExist(tb testing.TB, name string) {
	tb.Helper()
	_, err := os.Stat(name)
	if err == nil {
		tb.Errorf("a file exists unexpectedly: %s", name)
		return
	}
	if !os.IsNotExist(err) {
		tb.Errorf("unexpected error, want fs.ErrNotExist: got=%s", err)
		return
	}
}
