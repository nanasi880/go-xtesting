package xtesting

import (
	"io"
	"testing"
)

// MustClose is close a `c`. if error, test will fatal.
func MustClose(t testing.TB, c io.Closer) {
	t.Helper()

	if err := c.Close(); err != nil {
		t.Fatal(err)
	}
}
