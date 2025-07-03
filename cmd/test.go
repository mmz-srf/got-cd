package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Test() {
	currentFeatureBranch := strings.TrimSuffix(getCurrentBranch(), "\n")
	if currentFeatureBranch == "main" || currentFeatureBranch == "test" {
		log.Fatalf("You are on the main or on the test branch. Please switch to a feature branch first.\n")
	}

	checkoutTestCmd := exec.Command("git", "checkout", "test")
	output, err := checkoutTestCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error checking out test branch: %v\nOutput: %s", err, output)
	}

	fmt.Printf("Merging feature branch %v into test\n", currentFeatureBranch)
	mergeTestCmd := exec.Command("git", "merge", currentFeatureBranch)
	output, err = mergeTestCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error merging feature branch into test: %v\nOutput: %s", err, output)
	}

	pushTestCmd := exec.Command("git", "push", "origin", "test")
	output, err = pushTestCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error pushing test branch to origin: %v\nOutput: %s", err, output)
	}

	checkoutFeatureBranchCmd := exec.Command("git", "checkout", currentFeatureBranch)
	output, err = checkoutFeatureBranchCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error checking out feature branch: %v\nOutput: %s", err, output)
	}
}
