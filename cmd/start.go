package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/michizubi-SRF/got-cd/internal/helper"
)

func Start(branchName string, isVerbose bool) {

	branchName = helper.ReplaceSpacesWithDashes(branchName)

	if helper.GetCurrentBranch() != "main" && helper.GetCurrentBranch() != "master" {
		log.Fatal(helper.FormatMessage("You are not on the main/master branch. Please switch to the main/master branch before starting a new feature branch.", "warning"))
	}

	if isVerbose {
		fmt.Print(helper.FormatMessage("git pull", "verbose"))
	}
	pullCmd := exec.Command("git", "pull")
	output, err := pullCmd.CombinedOutput()
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error pulling latest changes: %v\n%s", "error"), err, output)
	}

	fmt.Printf(helper.FormatMessage("Creating new branch: %s", "info"), branchName)

	if isVerbose {
		fmt.Printf(helper.FormatMessage("git checkout -b %s", "verbose"), branchName)
	}
	checkoutCmd := exec.Command("git", "checkout", "-b", branchName)
	output, err = checkoutCmd.CombinedOutput()
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error creating new branch: %v\n%s", "error"), err, output)
	}

	if isVerbose {
		fmt.Printf(helper.FormatMessage("git push --set-upstream origin %s", "verbose"), branchName)
	}
	pushCmd := exec.Command("git", "push", "--set-upstream", "origin", branchName)
	output, err = pushCmd.CombinedOutput()
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error pushing new branch to origin: %v\n%s", "error"), err, output)
	}
}

func StartAsk(debugEnabled bool) {
	branchName, err := helper.AskForInput("Enter the name of the new feature branch:")
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error reading input: %v", "error"), err)
	}

	if branchName == "" {
		log.Fatal(helper.FormatMessage("Branch name cannot be empty.", "error"))
	}

	Start(branchName, debugEnabled)
}
