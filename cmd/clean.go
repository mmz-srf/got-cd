package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"slices"
	"strings"
)

func Clean() {
	remoteBranches, err := getRemoteBranches()
	if err != nil {
		log.Fatal("Error getting remote branches:", err)
	}

	localBranches, err := getLocalBranches()
	if err != nil {
		log.Fatal("Error getting local branches:", err)
	}

	for _, localBranch := range localBranches {
		if !slices.Contains(remoteBranches, localBranch) {
			fmt.Println("Found local branch not available on remote:", localBranch)
			fmt.Println("Do you want to delete this branch? (y/n)")
			var response string
			fmt.Scan(&response)
			if response == "y" {
				cmd := exec.Command("git", "branch", "-d", localBranch)
				if err := cmd.Run(); err != nil {
					log.Fatalf("Error deleting local branch %s: %v\n", localBranch, err)
				}
			}
		}
	}
}

// Get all remote branches
func getRemoteBranches() ([]string, error) {
	cmd := exec.Command("git", "fetch", "--prune")
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("error fetching remote branches: %w", err)
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

func getLocalBranches() ([]string, error) {
	cmd := exec.Command("git", "branch")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting local branches: %w", err)
	}

	branches := strings.Split(string(output), "\n")
	var localBranches []string
	for _, localBranch := range branches {
		localBranchCleaned := strings.Replace(string(localBranch), "* ", "", 1)
		localBranchTrimmed := strings.TrimSpace(localBranchCleaned)
		localBranches = append(localBranches, localBranchTrimmed)
	}

	return localBranches, nil
}
