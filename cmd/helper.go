package cmd

import (
	"log"
	"os/exec"
	"strings"
)

func getCurrentBranch() string {
	currentBranchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	currentBranchOutput, err := currentBranchCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("error getting current branch: %v\nOutput: %s", err, currentBranchOutput)
	}

	return string(currentBranchOutput)
}

func getCurrentFeatureBranch() string {
	currentFeatureBranch := strings.TrimSuffix(getCurrentBranch(), "\n")

	return currentFeatureBranch
}

func getCurrentRepoName() string {
	currentRepoPathOutput := exec.Command("git", "rev-parse", "--show-toplevel")
	currentRepoPath, err := currentRepoPathOutput.CombinedOutput()
	if err != nil {
		log.Fatalf("Error getting current repository path: %v\nOutput: %s", err, currentRepoPath)
	}
	currentRepoName := exec.Command("basename", string(currentRepoPath))
	repoResult, err := currentRepoName.CombinedOutput()
	if err != nil {
		log.Fatalf("Error getting current repository name: %v\nOutput: %s", err, repoResult)
	}
	repoName := strings.TrimSpace(string(repoResult))

	return repoName
}
