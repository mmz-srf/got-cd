package cmd

import (
	"fmt"
	"log"

	"github.com/google/go-github/v53/github"
)

func Status() {
	currentFeatureBranch := getCurrentFeatureBranch()
	repoName := getCurrentRepoName()

	if currentFeatureBranch == "main" || currentFeatureBranch == "test" {
		log.Fatalf("You are on the main or on the test branch. Please switch to a feature branch first.\n")
	}

	ctx, client := Authenticate()

	config := readConfigFile()
	githubOrganization := config.GithubOrganization
	existingOpenPRs, _, err := client.PullRequests.List(ctx, githubOrganization, string(repoName), &github.PullRequestListOptions{
		Head: fmt.Sprint(currentFeatureBranch),
	})
	if err != nil {
		log.Fatalf("Error checking for existing PR: %v\n", err)
	}

	for _, existingPR := range existingOpenPRs {
		if existingPR.GetTitle() == fmt.Sprint(currentFeatureBranch) {
			fmt.Printf("State of the PR is: %v \n", existingPR.GetState())
			reviews, _, err := client.PullRequests.ListReviews(ctx, githubOrganization, string(repoName), *existingPR.Number, &github.ListOptions{})
			if err != nil {
				log.Fatalf("Error getting reviews for PR: %v\n", err)
			}
			for _, review := range reviews {
				fmt.Printf("Review Comment by %s: %s\n", review.GetUser().GetLogin(), review.GetBody())
			}

			break
		}
	}
}
