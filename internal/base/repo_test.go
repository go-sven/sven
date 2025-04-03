package base

import (
	"context"
	"fmt"
	"path"
	"testing"
)

func TestRepo(t *testing.T) {
	ctx := context.Background()
	layout := "https://gitee.com/gosven/sven-layout.git"
	repo := NewRepo(layout, "")
	t.Log(repo.Path()) // ~/.sven/repo/layout@main
	err := repo.Clone(ctx)
	if err != nil {
		t.Error(err)
	}

	filename := path.Join(repo.Path(), "go.mod") ///Users/a1/.sven/repo/sven-layout@main/go.mod
	t.Log("filename:", filename)                 //github/go-sven/sven-layout

	fmt.Println("done")

}
