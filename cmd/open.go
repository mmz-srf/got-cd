package cmd

import (
	"fmt"
	"log"

	"github.com/michizubi-SRF/got-cd/internal/helper"
	"github.com/pkg/browser"
)

func Open() {
	currentFeatureBranch := helper.GetCurrentFeatureBranch()
	if currentFeatureBranch == "" {
		log.Fatal(helper.FormatMessage("No feature branch is currently checked out. Please switch to a feature branch before opening it in the browser.", "warning"))
	}

	repoName := helper.GetCurrentRepoName()
	if repoName == "" {
		log.Fatal(helper.FormatMessage("Could not determine the repository name. Please ensure you are in a git repository.", "error"))
	}

	repoUrl, err := helper.GetRemoteUrl()
	if err != nil {
		log.Fatal(helper.FormatMessage("Could not determine the remote URL. Please ensure you have a remote set up for this repository.", "error"))
	}

	fmt.Printf(helper.FormatMessage("Opening feature branch in the browser for repository : "+repoUrl, "info"))

	url := fmt.Sprintf("%s/tree/%s", repoUrl, currentFeatureBranch)

	err = browser.OpenURL(url)
	if err != nil {
		log.Fatal(helper.FormatMessage(fmt.Sprintf("Failed to open browser: %v", err), "error"))
	}

}
