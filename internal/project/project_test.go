package project

import (
	"os"
	"testing"
)

func TestProject(t *testing.T) {
	wd, _ := os.Getwd()
	t.Log("dir:", getProjectPlaceDir("demo", wd))
}
