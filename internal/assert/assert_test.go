package assert_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/koron-go/daemonic/internal/assert"
)

func TestEqual(t *testing.T) {
	assert.Equal(t, "foobar", "foobar")

	t.Run("failed", func(t *testing.T) {
		t2 := &testing.T{}
		assert.Equal(t2, "foo", "bar")
		if !t2.Failed() {
			t.Error("assert.Equal succeeded, unexpectedly")
		}
	})
}

func TestIsNotExist(t *testing.T) {
	tmpdir := t.TempDir()
	assert.IsNotExist(t, filepath.Join(tmpdir, "notexist.txt"))

	t.Run("failed", func(t *testing.T) {
		tmpdir := t.TempDir()
		name := filepath.Join(tmpdir, "exist.txt")
		err := os.WriteFile(name, []byte("hello"), 0666)
		if err != nil {
			t.Fatal(err)
		}

		t2 := &testing.T{}
		assert.IsNotExist(t2, name)
		if !t2.Failed() {
			t.Error("assert.IsNotExist succeeded, unexpectedly")
		}
	})
}
