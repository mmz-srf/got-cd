package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/google/go-github/v73/github"
	"github.com/michizubi-SRF/got-cd/internal/helper"
)

func Preview() {
	currentFeatureBranch := helper.GetCurrentFeatureBranch()
	repoName := helper.GetCurrentRepoName()

	if currentFeatureBranch == "main" || currentFeatureBranch == "test" {
		log.Fatal(helper.FormatMessage("You are on the main or on the test branch. Please switch to a feature branch first.\n", "warning"))
	}

	fmt.Printf(helper.FormatMessage("Creating pull request from feature branch %v to main\n", "info"), currentFeatureBranch)
	ctx, client := helper.Authenticate()

	config := helper.ReadConfigFile()
	githubOrganization := config.GithubOrganization
	existingOpenPRs, _, err := client.PullRequests.List(ctx, githubOrganization, string(repoName), &github.PullRequestListOptions{
		State: "open",
		Head:  fmt.Sprint(currentFeatureBranch),
	})
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error checking for existing PR: %v\n", "error"), err)
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
		fmt.Println(helper.FormatMessage("PR already exists, do you want to open it in the browser? (y/n)", "warning"))
		var response string
		fmt.Scan(&response)
		if response == "y" {
			_ = exec.Command("open", existingOpenPR.GetHTMLURL()).Start()
		}
	} else {

		newPR := &github.NewPullRequest{
			Title: github.String(fmt.Sprint(currentFeatureBranch)),
			Head:  github.String(currentFeatureBranch),
			Base:  github.String("main"),
		}

		pr, _, err := client.PullRequests.Create(ctx, "michizubi-SRF", string(repoName), newPR)
		if err != nil {
			log.Fatalf(helper.FormatMessage("Error creating pull request: %v\n", "error"), err)
		}
		fmt.Println(helper.FormatMessage("Do you want to open the PR in your browser? (y/n)", "info"))
		var response string
		fmt.Scan(&response)
		if response == "y" {
			err = exec.Command("open", pr.GetHTMLURL()).Start()
			if err != nil {
				log.Fatalf(helper.FormatMessage("Error opening PR in browser: %v\n", "error"), err)
			}
		}
	}

}
