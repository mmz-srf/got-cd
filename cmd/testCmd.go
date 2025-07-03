package cmd

import (
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Merge feature branch into test",
	Long:  `Merge the current feature branch into the test branch in the git repository.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		Test()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
