package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	GithubAccessToken  string `json:"github_access_token"`
	GithubOrganization string `json:"github_organization"`
}

func login() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error finding home directory:", err)
		os.Exit(1)
	}
	configDir := filepath.Join(homeDir, ".got-cd")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		fmt.Println("Error creating config directory:", err)
		os.Exit(1)
	}
	configPath := filepath.Join(configDir, "config.json")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to got-cd! Please log in to GitHub to continue.")
	fmt.Println("You will need a GitHub access toke (classic) with the 'repo' scope to use this tool.")
	fmt.Println("If you don't have a token yet, you can create one at https://github.com/settings/tokens\n\n")

	fmt.Print("Enter your GitHub access token: ")
	token, _ := reader.ReadString('\n')
	token = strings.TrimSpace(token)

	fmt.Print("Enter your GitHub organization: ")
	org, _ := reader.ReadString('\n')
	org = strings.TrimSpace(org)

	config := Config{
		GithubAccessToken:  token,
		GithubOrganization: org,
	}
	file, err := os.Create(configPath)
	if err != nil {
		fmt.Println("Error creating config file:", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(config); err != nil {
		fmt.Println("Error writing config:", err)
		os.Exit(1)
	}

	fmt.Printf("Config saved to %s\n", configPath)
}
