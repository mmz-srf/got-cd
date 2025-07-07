package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	GithubAccessToken  string `json:"github_access_token"`
	GithubOrganization string `json:"github_organization"`
}

func ReadConfigFile() Config {
	usersHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting user home directory: %v\n", err)
	}
	configFile := usersHomeDir + "/.got-cd/config.json"
	fileBytes, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error opening config file: %v\n", err)
	}

	var config Config
	json.NewDecoder((bytes.NewBuffer(fileBytes))).Decode(&config)

	return config

}
