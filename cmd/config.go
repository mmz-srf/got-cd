package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	GithubAccessToken  string `json:"github_access_token"`
	GithubOrganization string `json:"github_organization"`
}

func readConfigFile() Config {
	currentWorkingDirectory := getCurrentWorkingDirectory()
	configFile := currentWorkingDirectory + "/config.json"

	fileBytes, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error opening config file: %v\n", err)
	}

	var config Config
	json.NewDecoder((bytes.NewBuffer(fileBytes))).Decode(&config)

	return config

}
