package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/michizubi-SRF/got-cd/internal/helper"
)

func Finish() {
	currentFeatureBranch := helper.GetCurrentFeatureBranch()
	if currentFeatureBranch == "main" {
		fmt.Println("You are already on the main branch. Switch to your feature branch first.")
	}

	mergeCmd := exec.Command("git", "checkout", "main")
	if err := mergeCmd.Run(); err != nil {
		log.Fatalf("Error switching to main branch: %v\n", err)
	}

	mergeCmd = exec.Command("git", "merge", currentFeatureBranch)
	if err := mergeCmd.Run(); err != nil {
		log.Fatalf("Error merging feature branch into main: %v\n", err)
	}

	pushCmd := exec.Command("git", "push", "origin", "main")
	if err := pushCmd.Run(); err != nil {
		log.Fatalf("Error pushing changes to main branch: %v\n", err)
	}

	fmt.Println("Do you want to delete your feature branch? (y/n)")
	var response string
	fmt.Scan(&response)
	if response == "y" {
		deleteBranchCmd := exec.Command("git", "branch", "-d", currentFeatureBranch)
		if err := deleteBranchCmd.Run(); err != nil {
			log.Fatalf("Error deleting feature branch: %v\n", err)
		}
		fmt.Printf("Feature branch %s deleted.\n", currentFeatureBranch)
	}
}
