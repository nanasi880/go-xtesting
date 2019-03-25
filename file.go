package xtesting

import (
	"os"
	"testing"
)

// MustOpen is open file. if error, test will fatal.
func MustOpen(t testing.TB, name string) *os.File {
	t.Helper()

	f, err := os.Open(name)
	if err != nil {
		t.Fatal(err)
	}

	return f
}

// MustOpenFile is open file with flag. if error, test will fatal.
func MustOpenFile(t testing.TB, name string, flag int, perm os.FileMode) *os.File {
	t.Helper()

	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		t.Fatal(err)
	}

	return f
}
