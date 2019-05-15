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
	if !strings.HasPrefix(abs, "/tmp") || !strings.HasSuffix(abs, "gopwd.test") {
		t.Errorf("Abs() returned %s (expected `/tmp/*/gopwd.test`)", abs)
	}
}

func TestName(t *testing.T) {
	name, err := Name()
	if err != nil {
		t.Errorf("Name() returned error: %s", err)
	}
	// golang tests are running in `/tmp/go***/` dir
	if name != "gopwd.test" {
		t.Errorf("Name() returned %s (expected `gopwd.test`)", name)
	}
}
