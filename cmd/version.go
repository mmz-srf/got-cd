package cmd

import (
	_ "embed"
)

const currentVersionFallback = "dev" // First version with the version file

//go:embed VERSION
var currentVersion string

func GetVersion() string {
	if currentVersion == "" {
		return currentVersionFallback
	}
	return currentVersion
}
