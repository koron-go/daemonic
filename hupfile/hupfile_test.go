package hupfile_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"testing"
	"time"

	"github.com/koron-go/daemonic/hupfile"
	"github.com/koron-go/daemonic/internal/assert"
)

func TestHupfile(t *testing.T) {
	//if runtime.GOOS == "windows" {
	//	t.Skip("no SIGHUP support on Windows")
	//}
	tmpdir := t.TempDir()

	name := filepath.Join(tmpdir, "access.log")
	name0 := filepath.Join(tmpdir, "access.log.0")

	w, err := hupfile.New(name)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Fprint(w, "foo\n")

	// Reopen a file
	switch runtime.GOOS {
	case "windows":
		if err := w.Reopen(); err != nil {
			t.Error(err)
		}
		if err := os.Rename(name, name0); err != nil {
			t.Error(err)
		}

	default:
		if err := os.Rename(name, name0); err != nil {
			t.Error(err)
		}
		process, err := os.FindProcess(os.Getpid())
		if err != nil {
			t.Fatal(err)
		}
		if err := process.Signal(syscall.SIGHUP); err != nil {
			t.Error(err)
		}
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Fprint(w, "bar\n")

	if err := w.Close(); err != nil {
		t.Error(err)
	}

	got, err := os.ReadFile(name)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "bar\n", string(got))

	got0, err := os.ReadFile(name0)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "foo\n", string(got0))
}
