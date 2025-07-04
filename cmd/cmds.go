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

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Merge feature branch into test",
	Long:  `Merge the current feature branch into the test branch in the git repository.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		Test()
	},
}

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Create a pull request from feature branch to main",
	Long:  `Create a pull request from the current feature branch to the main branch in the git repository.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		Preview()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(reviewCmd)
}
