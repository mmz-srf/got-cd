package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func GetCurrentBranch() string {
	currentBranchCmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")
	currentBranchOutput, err := currentBranchCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("error getting current branch: %v\nOutput: %s", err, currentBranchOutput)
	}

	return string(currentBranchOutput)
}

func GetCurrentFeatureBranch() string {
	currentFeatureBranch := strings.TrimSuffix(GetCurrentBranch(), "\n")

	return currentFeatureBranch
}

func GetCurrentRepoName() string {
	currentRepoPathOutput := exec.Command("git", "rev-parse", "--show-toplevel")
	currentRepoPath, err := currentRepoPathOutput.CombinedOutput()
	if err != nil {
		log.Fatalf("Error getting current repository path: %v\nOutput: %s", err, currentRepoPath)
	}
	currentRepoName := exec.Command("basename", string(currentRepoPath))
	repoResult, err := currentRepoName.CombinedOutput()
	if err != nil {
		log.Fatalf("Error getting current repository name: %v\nOutput: %s", err, repoResult)
	}
	repoName := strings.TrimSpace(string(repoResult))

	return repoName
}

func GetCurrentWorkingDirectory() string {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
	}
	return currentWorkingDirectory
}

func FormatMessage(message string, level string) string {
	switch level {
	case "info":
		return fmt.Sprintf("\033[0;32m%s\033[0m\n", message) // Green
	case "error":
		return fmt.Sprintf("\033[0;31m%s\033[0m\n", message) // Red
	case "warning":
		return fmt.Sprintf("\033[0;33m%s\033[0m\n", message) // Yellow
	default:
		return message
	}
}

func GetDevBranch() string {
	branches, err := GetRemoteBranches()
	var devBranch string
	if err != nil {
		log.Fatalf(FormatMessage("Error getting local branches: %v", "error"), err)
	}

	reDevBranch := regexp.MustCompile(`^origin/(test|dev)-.*`)
	for _, branch := range branches {
		if reDevBranch.MatchString(branch) {
			devBranch = strings.Replace(branch, "origin/", "", 1)
		}
	}

	return devBranch

}

func GetLocalBranches() ([]string, error) {
	cmd := exec.Command("git", "branch")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting local branches: %w", err)
	}

	branches := strings.Split(string(output), "\n")
	var localBranches []string
	for _, localBranch := range branches {
		localBranchCleaned := strings.Replace(string(localBranch), "* ", "", 1)
		localBranchTrimmed := strings.TrimSpace(localBranchCleaned)
		localBranches = append(localBranches, localBranchTrimmed)
	}

	return localBranches, nil
}

func GetRemoteBranches() ([]string, error) {
	cmd := exec.Command("git", "branch", "-r")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting local branches: %w", err)
	}

	branches := strings.Split(string(output), "\n")
	var localBranches []string
	for _, localBranch := range branches {
		localBranchCleaned := strings.Replace(string(localBranch), "* ", "", 1)
		localBranchTrimmed := strings.TrimSpace(localBranchCleaned)
		localBranches = append(localBranches, localBranchTrimmed)
	}

	return localBranches, nil
}

func GetRemoteUrl() (string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin") // TODO: The remote name "origin" is hardcoded, consider making it configurable
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting remote URL: %w", err)
	}

	repoUrl := strings.TrimSpace(string(output))

	if len(repoUrl) > 0 && strings.HasPrefix(repoUrl, "git@") {
		// Pattern: git@hostname:username/repository.git
		// Example: git@github.com:mmz-srf/got-cd.git
		gitUrlPattern := regexp.MustCompile(`^git@([^:]+):(.+)\.git$`)
		if matches := gitUrlPattern.FindStringSubmatch(repoUrl); matches != nil {
			repoUrl = fmt.Sprintf("https://%s/%s", matches[1], matches[2])
		}
	}

	return strings.TrimSpace(repoUrl), nil
}

func AskForInput(prompt string) (string, error) {
	fmt.Print(prompt + " ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text, err
}

func ReplaceSpacesWithDashes(input string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(input, "-")
}
