package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/google/go-github/v53/github"
)

func Preview() {
	currentFeatureBranch := strings.TrimSuffix(getCurrentBranch(), "\n")
	if currentFeatureBranch == "main" || currentFeatureBranch == "test" {
		log.Fatalf("You are on the main or on the test branch. Please switch to a feature branch first.\n")
	}

	fmt.Printf("Creating pull request from feature branch %v to main\n", currentFeatureBranch)
	ctx, client := Authenticate()

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

	newPR := &github.NewPullRequest{
		Title: github.String(fmt.Sprintf("Feature: %s", currentFeatureBranch)),
		Head:  github.String(currentFeatureBranch),
		Base:  github.String("main"),
	}

	pr, _, err := client.PullRequests.Create(ctx, "michizubi-SRF", string(repoName), newPR)
	if err != nil {
		log.Fatalf("Error creating pull request: %v\n", err)
	}
	fmt.Printf("PR created: %s\n", pr.GetHTMLURL())
}
