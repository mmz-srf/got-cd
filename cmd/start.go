package cmd

import (
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [branch-name]",
	Short: "Start a new feature branch",
	Long:  `Start a new feature branch by creating a new branch in the git repository.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Start(args[0])
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
