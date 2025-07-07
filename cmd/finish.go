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
		return
	}

	mergeCmd := exec.Command("git", "checkout", "main")
	if err := mergeCmd.Run(); err != nil {
		fmt.Printf("Error switching to main branch: %v\n", err)
		return
	}

	mergeCmd = exec.Command("git", "merge", currentFeatureBranch)
	if err := mergeCmd.Run(); err != nil {
		fmt.Printf("Error merging feature branch into main: %v\n", err)
		return
	}

	pushCmd := exec.Command("git", "push", "origin", "main")
	if err := pushCmd.Run(); err != nil {
		log.Fatalf("Error pushing changes to main branch: %v\n", err)
	}
}
