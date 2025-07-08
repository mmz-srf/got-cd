package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/michizubi-SRF/got-cd/internal/helper"
)

func Start(branchName string) {

	if helper.GetCurrentBranch() != "main\n" {
		log.Fatal(helper.FormatMessage("You are not on the main branch. Please switch to the main branch before starting a new feature branch.", "warning"))
	}

	fmt.Printf(helper.FormatMessage("Creating new branch: %s", "info"), branchName)

	checkoutCmd := exec.Command("git", "checkout", "-b", branchName)
	output, err := checkoutCmd.CombinedOutput()
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error creating new branch: %v\n%s", "error"), err, output)
	}

	pushCmd := exec.Command("git", "push", "--set-upstream", "origin", branchName)
	output, err = pushCmd.CombinedOutput()
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error pushing new branch to origin: %v\n%s", "error"), err, output)
	}
}
