package helper

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	GithubAccessToken  string `json:"github_access_token"`
	GithubOrganization string `json:"github_organization"`
}

func ReadConfigFile() (Config, error) {
	usersHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf(FormatMessage("Error getting user home directory: %v\n", "error"), err)
	}
	configFile := usersHomeDir + "/.got-cd/config.json"
	fileBytes, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}

	var config Config
	json.NewDecoder((bytes.NewBuffer(fileBytes))).Decode(&config)

	return config, nil

}
