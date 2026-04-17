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
	currentBranch := strings.TrimSuffix(string(currentBranchOutput), "\n")
	if err != nil {
		log.Fatalf("error getting current branch: %v\nOutput: %s", err, currentBranchOutput)
	}

	return string(currentBranch)
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
	case "verbose":
		return fmt.Sprintf("\033[0;34m%s\033[0m\n", message) // Blue
	default:
		return message
	}
}

func GetGitRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not get repository root: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func ReadGitCdConfig(key string) (string, error) {
	repoRoot, err := GetGitRoot()
	if err != nil {
		return "", fmt.Errorf("could not get repository root: %w", err)
	}

	configPath := repoRoot + "/.gitcd"
	file, err := os.Open(configPath)
	if err != nil {
		return "", fmt.Errorf("could not open .gitcd config file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Support both "key: value" and "key=value" formats
		var configKey, configValue string
		if strings.Contains(line, ": ") {
			parts := strings.SplitN(line, ": ", 2)
			if len(parts) == 2 {
				configKey = strings.TrimSpace(parts[0])
				configValue = strings.TrimSpace(parts[1])
			}
		}

		if configKey == key {
			return configValue, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading .gitcd config file: %w", err)
	}

	return "", fmt.Errorf("key '%s' not found in .gitcd config", key)
}

func GetDevBranchPrefix() string {
	prefix, err := ReadGitCdConfig("test")
	if err != nil {
		log.Printf(FormatMessage("Warning: Could not read dev-branch-prefix from config, using default", "warning"))
		return "dev"
	}

	if prefix == "" {
		return "dev"
	}

	return prefix
}

func GetDevBranch() string {
	branches, err := GetRemoteBranches()
	var devBranch string
	if err != nil {
		log.Fatalf(FormatMessage("Error getting local branches: %v", "error"), err)
	}

	prefix := GetDevBranchPrefix()
	reDevBranch := regexp.MustCompile(fmt.Sprintf(`^origin/(%s)-.*`, prefix))
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

func GetNameOfDefaultBranch() string {
	cmd := exec.Command("bash", "-c", "git remote show origin|awk '/HEAD branch|Hauptbranch/ {print $NF}'")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error getting default branch name: %v\nOutput: %s", err, output)
	}
	defaultBranch := strings.TrimSpace(string(output))
	return defaultBranch
}
