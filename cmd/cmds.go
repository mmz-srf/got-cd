package cmd

import (
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [branch-name]",
	Short: "Start a new feature branch",
	Long:  `Start a new feature branch by creating a new branch in the git repository.`,
	//Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			StartAsk()
			return
		}
		if len(args) == 1 {
			Start(args[0])
			return
		}
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

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Create a new release",
	Long:  `Create a new release by pushing a new tag based on version.txt.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		Release()
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status of an open pull request",
	Long:  `Get the status of an open pull request from the current feature branch.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		Status()
	},
}

var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Merge the feature branch into main",
	Long:  `Merge the feature branch into main.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		Finish()
	},
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up local branches",
	Long:  `Clean up local branches that are not present on remote anymore.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		Clean()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of got-cd",
	Long:  `Print the version of got-cd.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		println(GetVersion())
	},
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open the current feature branch in the browser",
	Long:  `Open the current feature branch in the browser.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		Open()
	},
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to GitHub",
	Long:  `Login to GitHub by providing your access token.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(reviewCmd)
	rootCmd.AddCommand(releaseCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(finishCmd)
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(loginCmd)
}
