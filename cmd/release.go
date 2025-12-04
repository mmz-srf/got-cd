package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/michizubi-SRF/got-cd/internal/helper"
)

func Release(isVerbose bool, isShortTag bool) {
	currentFeatureBranch := strings.TrimSuffix(helper.GetCurrentBranch(), "\n")
	if currentFeatureBranch != "main" && currentFeatureBranch != "master" {
		println(helper.FormatMessage("You are not on the main/master branch. Please switch to the main branch before releasing.", "warning"))
		fmt.Print(helper.FormatMessage("Do you want to switch to the main/master branch? (y/n)", "info"))
		var response string
		fmt.Scan(&response)
		if response == "y" {
			if isVerbose {
				fmt.Print(helper.FormatMessage("git checkout main/master", "verbose"))
			}
			defaultBranch := helper.GetNameOfDefaultBranch()
			switchCmd, err := exec.Command("git", "checkout", defaultBranch).CombinedOutput()
			if err != nil {
				log.Fatalf(helper.FormatMessage("Error switching to main/master branch: %v\nOutput: %s", "error"), err, switchCmd)
			}
		}
	}

	currentWorkingDirectory := helper.GetCurrentWorkingDirectory()
	var releaseVersion string
	if _, err := os.Stat(currentWorkingDirectory + "/version.txt"); err == nil {
		releaseVersionFile, err := os.Open(currentWorkingDirectory + "/version.txt")
		if err != nil {
			log.Fatalf(helper.FormatMessage("Error opening version.txt: %v", "error"), err)
		}
		defer releaseVersionFile.Close()

		releaseVersionFromFile, err := io.ReadAll(releaseVersionFile)
		if err != nil {
			log.Fatalf(helper.FormatMessage("Error reading version.txt: %v", "error"), err)
		}
		releaseVersion = string(releaseVersionFromFile)
	} else {
		var latestTag string
		if isVerbose {
			fmt.Print(helper.FormatMessage("git fetch --tags", "verbose"))
		}
		fetchTagsCmd, err := exec.Command("git", "fetch", "--tags").CombinedOutput()
		if err != nil {
			log.Fatalf(helper.FormatMessage("Error fetching tags: %v\n%s", "error"), err, fetchTagsCmd)
		}

		if isVerbose {
			fmt.Print(helper.FormatMessage("git describe --tags --abbrev=0", "verbose"))
		}
		latestTagCmd, err := exec.Command("git", "describe", "--tags", "--abbrev=0").CombinedOutput()
		if err != nil {
			log.Fatalf(helper.FormatMessage("Error getting latest tag: %v\n%s", "error"), err, latestTagCmd)
		}
		latestTag = strings.TrimSpace(string(latestTagCmd))
		fmt.Printf(helper.FormatMessage("Latest tag found: %s\n", "info"), latestTag)
		fmt.Print(helper.FormatMessage("version.txt not found. Please specify version to release:", "info"))

		var manualVersion string
		fmt.Scan(&manualVersion)
		releaseVersion = manualVersion
	}

	var versionTag string
	if isShortTag {
		versionTag = string(releaseVersion)
	} else {
		versionTag = "v" + string(releaseVersion)
	}

	versionTagTrimmed := strings.TrimSuffix(versionTag, "\n")

	fmt.Printf(helper.FormatMessage("Releasing version: %s", "info"), versionTagTrimmed)
	fmt.Println((helper.FormatMessage("What is this release about?", "info")))
	var releaseMessage string
	reader := bufio.NewReader(os.Stdin)
	releaseMessage, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error reading release message: %v", "error"), err)
	}
	releaseMessage = strings.TrimSpace(releaseMessage)

	if isVerbose {
		fmt.Printf(helper.FormatMessage("git tag -a -m \"%s\" %s", "verbose"), releaseMessage, versionTagTrimmed)
	}
	releaseCmd, err := exec.Command("git", "tag", "-a", "-m", releaseMessage, versionTagTrimmed).CombinedOutput()
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error creating release tag: %v\n%s", "error"), err, releaseCmd)
	}
	if isVerbose {
		fmt.Printf(helper.FormatMessage("git push origin %s", "verbose"), versionTagTrimmed)
	}
	tagPushCmd, err := exec.Command("git", "push", "origin", versionTagTrimmed).CombinedOutput()
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error pushing release tag to origin: %v\n%s", "error"), err, tagPushCmd)
	}

}
