package bootstrap

import (
	"os"
	"testing"
)

func TestInitialize(t *testing.T) {
	os.Setenv("TEST_MODE", "true")
	r := Initialize()
	if r == nil {
		t.Fatal("expected gin engine")
	}
}
