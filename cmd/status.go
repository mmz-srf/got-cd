package cmd

import (
	"fmt"
	"log"

	"github.com/google/go-github/v73/github"
	"github.com/michizubi-SRF/got-cd/internal/helper"
)

func Status() {
	currentFeatureBranch := helper.GetCurrentFeatureBranch()
	repoName := helper.GetCurrentRepoName()

	if currentFeatureBranch == "main" || currentFeatureBranch == "test" {
		log.Fatal(helper.FormatMessage("You are on the main or on the test branch. Please switch to a feature branch first.\n", "warning"))
	}

	ctx, client := helper.Authenticate()

	config, err := helper.ReadConfigFile()
	if err != nil {
		log.Fatalf(helper.FormatMessage("Status: Error opening config file: %v\n", "error"), err)
	}

	githubOrganization := config.GithubOrganization
	existingOpenPRs, _, err := client.PullRequests.List(ctx, githubOrganization, string(repoName), &github.PullRequestListOptions{
		Head: fmt.Sprint(currentFeatureBranch),
	})
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error checking for existing PR: %v\n", "error"), err)
	}

	if len(existingOpenPRs) == 0 {
		log.Fatalf(helper.FormatMessage("No open PR found for branch %s. Please create a PR first.\n", "warning"), currentFeatureBranch)
	}

	for _, existingPR := range existingOpenPRs {
		if existingPR.GetTitle() == fmt.Sprint(currentFeatureBranch) {
			fmt.Printf(helper.FormatMessage("State of the PR is: %v \n", "info"), existingPR.GetState())
			reviews, _, err := client.PullRequests.ListReviews(ctx, githubOrganization, string(repoName), *existingPR.Number, &github.ListOptions{})
			if err != nil {
				log.Fatalf(helper.FormatMessage("Error getting reviews for PR: %v\n", "error"), err)
			}
			for _, review := range reviews {
				fmt.Printf(helper.FormatMessage("Review Comment by %s: %s\n", "info"), review.GetUser().GetLogin(), review.GetBody())
			}

			break
		}
	}
}
