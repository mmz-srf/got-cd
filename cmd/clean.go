package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"slices"
	"strings"

	"github.com/michizubi-SRF/got-cd/internal/helper"
)

func Clean(isVerbose bool) {
	fmt.Println(helper.FormatMessage("Cleaning up local branches not available on remote", "info"))
	remoteBranches, err := getRemoteBranches(isVerbose)
	if err != nil {
		log.Fatal(helper.FormatMessage("Error getting remote branches:", "error"), err)
	}

	localBranches, err := helper.GetLocalBranches()
	if err != nil {
		log.Fatal(helper.FormatMessage("Error getting local branches:", "error"), err)
	}

	for _, localBranch := range localBranches {
		if !slices.Contains(remoteBranches, localBranch) {
			fmt.Println(helper.FormatMessage("Found local branch not available on remote:", "warning"), localBranch)
			fmt.Print(helper.FormatMessage("Do you want to delete this branch? (y/n): ", "warning"))
			var response string
			fmt.Scan(&response)
			if response == "y" {
				if isVerbose {
					fmt.Printf(helper.FormatMessage("git branch -D", "verbose"), localBranch)
				}
				cmd := exec.Command("git", "branch", "-D", localBranch)
				if err := cmd.Run(); err != nil {
					log.Fatalf(helper.FormatMessage("Error deleting local branch %s: %v\n", "error"), localBranch, err)
				}
			}
		}
	}
}

func getRemoteBranches(isVerbose bool) ([]string, error) {
	if isVerbose {
		fmt.Print(helper.FormatMessage("git fetch --prune", "verbose"))
	}
	cmd := exec.Command("git", "fetch", "--prune")
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("error fetching remote branches: %w", err)
	}
	if isVerbose {
		fmt.Print(helper.FormatMessage("git branch -r", "verbose"))
	}
	cmd = exec.Command("git", "branch", "-r")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting remote branches: %w", err)
	}

	branches := strings.Split(string(output), "\n")

	var remoteBranches []string
	for _, branch := range branches {
		brancheCleaned := strings.Replace(string(branch), "origin/", "", 1)
		branchTrimmed := strings.TrimSpace(brancheCleaned)
		remoteBranches = append(remoteBranches, branchTrimmed)
	}

	return remoteBranches, nil
}
