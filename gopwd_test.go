package gopwd

import (
	"strings"
	"testing"
)

func TestAbs(t *testing.T) {
	abs, err := Abs()
	if err != nil {
		t.Errorf("Abs() returned error: %s", err)
	}
	// golang tests are running in `/tmp/go***/` dir
	if !strings.HasSuffix(abs, "/gopwd") {
		t.Errorf("Abs() returned %s (expected `/*/*/gopwd`)", abs)
	}
}

func TestName(t *testing.T) {
	name, err := Name()
	if err != nil {
		t.Errorf("Name() returned error: %s", err)
	}
	if name == "" {
		t.Error("Name() returned empty")
	}
}
