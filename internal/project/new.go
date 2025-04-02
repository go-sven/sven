package project

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/go-sven/sven/internal/base"
	"os"
	"path"
)

// Project template
type Project struct {
	Name string
	Path string
}

// New create a project by layout repo
func (p *Project) New(ctx context.Context, dir string, layout string, branch string) error {
	to := path.Join(dir, p.Name)
	_, err := os.Stat(to)
	//如果项目已经存在
	if !os.IsNotExist(err) {
		fmt.Printf("🚫 %s already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "📂 是否覆盖该文件夹?",
			Help:    "删除该文件 然后重新创建新项目.",
		}
		if err = survey.AskOne(prompt, &override); err != nil {
			return err
		}
		if !override {
			return err
		}
		_ = os.RemoveAll(to)
	}
	fmt.Printf("Creating project %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)

	repo := base.NewRepo(layout, branch)
	//拉去layout 然后 copy到项目目录，然后替换名称
	err = repo.CopyTo(ctx, to, p.Path, []string{".git", ".github"})
	if err != nil {
		return err
	}
	base.Tree(to, dir)
	fmt.Printf("\n🍺 Project creation succeeded %s\n", color.GreenString(p.Name))
	return nil
}
