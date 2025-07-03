package cmd

import (
	"log"
	"os/exec"
)

func getCurrentBranch() string {
	currentBranchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	currentBranchOutput, err := currentBranchCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("error getting current branch: %v\nOutput: %s", err, currentBranchOutput)
	}

	return string(currentBranchOutput)
}
