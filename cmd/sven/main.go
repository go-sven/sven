package main

import (
	"log"
	"github.com/spf13/cobra"
	"github.com/go-sven/sven/cmd/sven/internal/project"
)



var rootCmd = &cobra.Command{
	Use:     "sven",
	Short:   "sven : The easy tool for gin",
	Long:    "sven : The easy tool for gin",
	Version: "v1.1.1",
}

func init()  {
	rootCmd.AddCommand(project.CmdNew)

}

func main()  {
	err := rootCmd.Execute()
	if err !=nil {
		log.Fatal(err)
	}
}
