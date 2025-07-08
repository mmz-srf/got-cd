package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/michizubi-SRF/got-cd/internal/helper"
)

func Test() {
	currentFeatureBranch := strings.TrimSuffix(helper.GetCurrentBranch(), "\n")
	if currentFeatureBranch == "main" || currentFeatureBranch == "test" {
		log.Fatal(helper.FormatMessage("You are on the main or on the test branch. Please switch to a feature branch first.\n", "warning"))
	}

	checkoutTestCmd := exec.Command("git", "checkout", "test")
	output, err := checkoutTestCmd.CombinedOutput()
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error checking out test branch: %v\n%s", "error"), err, output)
	}

	fmt.Printf("Merging feature branch %v into test\n", currentFeatureBranch)
	mergeTestCmd := exec.Command("git", "merge", currentFeatureBranch)
	output, err = mergeTestCmd.CombinedOutput()
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error merging feature branch into test: %v\n%s", "error"), err, output)
	}

	pushTestCmd := exec.Command("git", "push", "origin", "test")
	output, err = pushTestCmd.CombinedOutput()
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error pushing test branch to origin: %v\n%s", "error"), err, output)
	}

	checkoutFeatureBranchCmd := exec.Command("git", "checkout", currentFeatureBranch)
	output, err = checkoutFeatureBranchCmd.CombinedOutput()
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error checking out feature branch: %v\n%s", "error"), err, output)
	}
}
