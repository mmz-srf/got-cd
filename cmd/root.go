package cmd

import (
	"log"

	"github.com/michizubi-SRF/got-cd/internal/helper"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-cd",
	Short: "A wrapper for git commands",
	Long:  `git-cd is a command-line tool that simplifies the usage of git commands by providing a more user-friendly interface.`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf(helper.FormatMessage("Error executing command: %v", "error"), err)
	}
}
