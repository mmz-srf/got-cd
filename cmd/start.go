package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/michizubi-SRF/got-cd/internal/helper"
)

func Start(branchName string) {

	if helper.GetCurrentBranch() != "main\n" {
		log.Fatalf("You are not on the main branch. Please switch to the main branch before starting a new feature branch.")
	}

	fmt.Println("Creating new feature branch " + branchName)

	checkoutCmd := exec.Command("git", "checkout", "-b", branchName)
	output, err := checkoutCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error creating new branch: %v\nOutput: %s", err, output)
	}

	pushCmd := exec.Command("git", "push", "--set-upstream", "origin", branchName)
	output, err = pushCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error pushing new branch to origin: %v\nOutput: %s", err, output)
	}

}
