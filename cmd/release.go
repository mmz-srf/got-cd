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

func Release(isVerbose bool) {
	currentFeatureBranch := strings.TrimSuffix(helper.GetCurrentBranch(), "\n")
	if currentFeatureBranch != "main" {
		println(helper.FormatMessage("You are not on the main branch. Please switch to the main branch before releasing.", "warning"))
		fmt.Println(helper.FormatMessage("Do you want to switch to the main branch? (y/n)", "info"))
		var response string
		fmt.Scan(&response)
		if response == "y" {
			if isVerbose {
				fmt.Print(helper.FormatMessage("git checkout main", "verbose"))
			}
			switchCmd, err := exec.Command("git", "checkout", "main").CombinedOutput()
			if err != nil {
				log.Fatalf(helper.FormatMessage("Error switching to main branch: %v\nOutput: %s", "error"), err, switchCmd)
			}
		}
	}

	currentWorkingDirectory := helper.GetCurrentWorkingDirectory()
	releaseVersionFile, err := os.Open(currentWorkingDirectory + "/version.txt")
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error opening version.txt: %v", "error"), err)
	}
	defer releaseVersionFile.Close()

	releaseVersion, err := io.ReadAll(releaseVersionFile)
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error reading version.txt: %v", "error"), err)
	}

	fmt.Printf(helper.FormatMessage("Releasing version: %s", "info"), string(releaseVersion))
	fmt.Println((helper.FormatMessage("What is this release about?", "info")))
	var releaseMessage string
	reader := bufio.NewReader(os.Stdin)
	releaseMessage, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error reading release message: %v", "error"), err)
	}
	releaseMessage = strings.TrimSpace(releaseMessage)
	versionTag := "v" + string(releaseVersion)
	versionTagTrimmed := strings.TrimSuffix(versionTag, "\n")
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
