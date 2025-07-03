package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "got-cd",
	Short: "A wrapper for git commands",
	Long:  `got-cd is a command-line tool that simplifies the usage of git commands by providing a more user-friendly interface.`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
