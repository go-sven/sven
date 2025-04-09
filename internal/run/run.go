package run

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var CmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run project",
	Long:  "Run project. Example: kratos run",
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	var dir string
	cmdArgs, programArgs := splitArgs(cmd, args)
	if len(cmdArgs) > 0 {
		dir = cmdArgs[0]
	} else {
		dir = "./cmd/api/"
	}
	fd := exec.Command("go", append([]string{"run", "."}, programArgs...)...)
	fd.Stdout = os.Stdout
	fd.Stderr = os.Stderr
	fd.Dir = dir
	if err := fd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err.Error())
		return
	}

}

func splitArgs(cmd *cobra.Command, args []string) (cmdArgs, programArgs []string) {
	dashAt := cmd.ArgsLenAtDash()
	if dashAt >= 0 {
		return args[:dashAt], args[dashAt:]
	}
	return args, []string{}
}
