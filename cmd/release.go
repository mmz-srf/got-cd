package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Release() {
	currentFeatureBranch := strings.TrimSuffix(getCurrentBranch(), "\n")
	if currentFeatureBranch != "main" {
		println("You are not on the main branch. Please switch to the main branch before releasing.")
		fmt.Println("Do you want to switch to the main branch? (y/n)")
		var response string
		fmt.Scan(&response)
		if response == "y" {
			switchCmd, err := exec.Command("git", "checkout", "main").CombinedOutput()
			if err != nil {
				log.Fatalf("Error switching to main branch: %v\nOutput: %s", err, switchCmd)
			}
		}
	}

	currentWorkingDirectory := getCurrentWorkingDirectory()
	releaseVersionFile, err := os.Open(currentWorkingDirectory + "/version.txt")
	if err != nil {
		log.Fatalf("Error opening version.txt: %v", err)
	}
	defer releaseVersionFile.Close()

	releaseVersion, err := io.ReadAll(releaseVersionFile)
	if err != nil {
		log.Fatalf("Error reading version.txt: %v", err)
	}

	fmt.Println("Releasing version:", string(releaseVersion))
	fmt.Println(("What is this release about?"))
	var releaseMessage string
	reader := bufio.NewReader(os.Stdin)
	releaseMessage, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading release message: %v", err)
	}
	releaseMessage = strings.TrimSpace(releaseMessage)
	versionTag := "v" + string(releaseVersion)
	releaseCmd, err := exec.Command("git", "tag", "-a", versionTag, "-m", releaseMessage).CombinedOutput()
	if err != nil {
		log.Fatalf("Error creating release tag: %v\nOutput: %s", err, releaseCmd)
	}
	tagPushCmd, err := exec.Command("git", "push", "origin", versionTag).CombinedOutput()
	if err != nil {
		log.Fatalf("Error pushing release tag to origin: %v\nOutput: %s", err, tagPushCmd)
	}

}
