package main

import (
	"github.com/go-sven/sven/internal/project"
	"github.com/go-sven/sven/internal/run"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:     "sven",
	Short:   "sven: sven is best",
	Long:    `sven: new gin project tool `,
	Version: "1.0.1",
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(run.CmdRun)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
