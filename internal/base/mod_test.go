package base

import (
	"testing"
)

func TestMod(t *testing.T) {
	path, err := ModuleVersion("github.com/go-sven/sven/v2")
	if err != nil {
		t.Error(err)
	}
	t.Log(path)
}
