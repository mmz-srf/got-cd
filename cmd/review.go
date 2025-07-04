package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/google/go-github/v53/github"
)

func Preview() {
	currentFeatureBranch := getCurrentFeatureBranch()
	repoName := getCurrentRepoName()

	if currentFeatureBranch == "main" || currentFeatureBranch == "test" {
		log.Fatalf("You are on the main or on the test branch. Please switch to a feature branch first.\n")
	}

	fmt.Printf("Creating pull request from feature branch %v to main\n", currentFeatureBranch)
	ctx, client := Authenticate()

	config := readConfigFile()
	githubOrganization := config.GithubOrganization
	existingOpenPRs, _, err := client.PullRequests.List(ctx, githubOrganization, string(repoName), &github.PullRequestListOptions{
		State: "open",
		Head:  fmt.Sprint(currentFeatureBranch),
	})
	if err != nil {
		log.Fatalf("Error checking for existing PR: %v\n", err)
	}
	var foundExistingPR bool
	var existingOpenPR *github.PullRequest
	for _, existingPR := range existingOpenPRs {
		if existingPR.GetTitle() == fmt.Sprint(currentFeatureBranch) {
			foundExistingPR = true
			existingOpenPR = existingPR
			break
		}
	}
	if foundExistingPR {
		fmt.Println("PR already exists, do you want to open it in the browser? (y/n)")
		var response string
		fmt.Scan(&response)
		if response == "y" {
			_ = exec.Command("open", existingOpenPR.GetHTMLURL()).Start()
		}
	} else {

		newPR := &github.NewPullRequest{
			Title: github.String(fmt.Sprintf("Feature: %s", currentFeatureBranch)),
			Head:  github.String(currentFeatureBranch),
			Base:  github.String("main"),
		}

		pr, _, err := client.PullRequests.Create(ctx, "michizubi-SRF", string(repoName), newPR)
		if err != nil {
			log.Fatalf("Error creating pull request: %v\n", err)
		}
		fmt.Println("Do you want to open the PR in your browser? (y/n)")
		var response string
		fmt.Scan(&response)
		if response == "y" {
			err = exec.Command("open", pr.GetHTMLURL()).Start()
			if err != nil {
				log.Fatalf("Error opening PR in browser: %v\n", err)
			}
		}
	}

}
