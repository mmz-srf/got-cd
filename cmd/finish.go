package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/michizubi-SRF/got-cd/internal/helper"
)

func Finish(isVerbose bool) {
	currentFeatureBranch := helper.GetCurrentFeatureBranch()
	if currentFeatureBranch == "main" || currentFeatureBranch == "master" {
		log.Fatal(helper.FormatMessage("You are on the main/master branch. Switch to your feature branch first.", "error"))
	}

	fmt.Printf(helper.FormatMessage("Merging feature branch %s into main\n", "info"), currentFeatureBranch)

	if isVerbose {
		fmt.Print(helper.FormatMessage("git checkout main/master", "verbose"))
	}
	defaultBranch := helper.GetNameOfDefaultBranch()
	mergeCmd := exec.Command("git", "checkout", defaultBranch)
	if err := mergeCmd.Run(); err != nil {
		log.Fatalf(helper.FormatMessage("Error switching to main/master branch: %v\n", "error"), err)
	}

	if isVerbose {
		fmt.Printf(helper.FormatMessage("git merge %s", "verbose"), currentFeatureBranch)
	}
	mergeCmd = exec.Command("git", "merge", currentFeatureBranch)
	if err := mergeCmd.Run(); err != nil {
		log.Fatalf(helper.FormatMessage("Error merging feature branch into main/master: %v\n", "error"), err)
	}

	if isVerbose {
		fmt.Print(helper.FormatMessage("git push origin main/master", "verbose"))
	}
	defaultBranch = helper.GetNameOfDefaultBranch()
	pushCmd := exec.Command("git", "push", "origin", defaultBranch)
	if err := pushCmd.Run(); err != nil {
		log.Fatalf(helper.FormatMessage("Error pushing changes to main/master branch: %v\n", "error"), err)
	}

	fmt.Print(helper.FormatMessage("Do you want to delete your feature branch? (y/n)", "warning"))
	var response string
	fmt.Scan(&response)
	if response == "y" {
		if isVerbose {
			fmt.Printf(helper.FormatMessage("git branch -d %s", "verbose"), currentFeatureBranch)
		}
		deleteBranchCmd := exec.Command("git", "branch", "-d", currentFeatureBranch)
		if err := deleteBranchCmd.Run(); err != nil {
			log.Fatalf(helper.FormatMessage("Error deleting feature branch: %v\n", "error"), err)
		}
		fmt.Printf(helper.FormatMessage("Feature branch %s deleted.\n", "info"), currentFeatureBranch)
	}
}
