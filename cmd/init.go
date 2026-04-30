package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/michizubi-SRF/got-cd/internal/helper"
)

type gitCdConfig struct {
	master              string
	feature             string
	test                string
	tag                 string
	versionType         string
	versionScheme       string
	extraReleaseCommand string
}

func Init() {
	repoRoot, err := helper.GetGitRoot()
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error: not inside a git repository: %v", "error"), err)
	}

	configPath := repoRoot + "/.gitcd"

	if _, err := os.Stat(configPath); err == nil {
		fmt.Print(helper.FormatMessage("A .gitcd config file already exists. Overwrite it? [y/N]: ", "warning"))
		var answer string
		fmt.Scanln(&answer)
		if strings.ToLower(strings.TrimSpace(answer)) != "y" {
			fmt.Println(helper.FormatMessage("Aborted.", "info"))
			return
		}
	}

	cfg := askForConfig()
	writeConfig(configPath, cfg)

	fmt.Printf(helper.FormatMessage("Created .gitcd config at %s", "info"), configPath)
}

func askForConfig() gitCdConfig {
	cfg := gitCdConfig{}

	cfg.master = askWithDefault("Master/main branch name", "main")
	cfg.feature = askWithDefault("Feature branch prefix", "feature/")
	cfg.test = askWithDefault("Test/develop branch name", "develop")
	cfg.tag = askWithDefault("Tag prefix", "v")
	cfg.versionType = askChoice(
		"Version type",
		[]string{"manual", "auto"},
		"manual",
	)
	cfg.versionScheme = askWithDefault("Version scheme (leave empty for none)", "")
	cfg.extraReleaseCommand = askWithDefault("Extra release command (leave empty for none)", "")

	return cfg
}

func askWithDefault(prompt, defaultValue string) string {
	if defaultValue != "" {
		fmt.Printf(helper.FormatMessage(prompt+" [%s]: ", "info"), defaultValue)
	} else {
		fmt.Print(helper.FormatMessage(prompt+": ", "info"))
	}

	input, err := helper.AskForInput("")
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error reading input: %v", "error"), err)
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue
	}
	return input
}

func askChoice(prompt string, choices []string, defaultChoice string) string {
	choiceStr := strings.Join(choices, "/")
	fmt.Printf(helper.FormatMessage(prompt+" [%s] (default: %s): ", "info"), choiceStr, defaultChoice)

	input, err := helper.AskForInput("")
	if err != nil {
		log.Fatalf(helper.FormatMessage("Error reading input: %v", "error"), err)
	}

	input = strings.TrimSpace(strings.ToLower(input))
	if input == "" {
		return defaultChoice
	}

	for _, c := range choices {
		if input == c {
			return input
		}
	}

	fmt.Printf(helper.FormatMessage("Invalid choice '%s', using default '%s'.", "warning"), input, defaultChoice)
	return defaultChoice
}

func nullOrValue(v string) string {
	if v == "" {
		return "null"
	}
	return v
}

func writeConfig(path string, cfg gitCdConfig) {
	lines := []string{
		fmt.Sprintf("extraReleaseCommand: %s", nullOrValue(cfg.extraReleaseCommand)),
		fmt.Sprintf("feature: %s", cfg.feature),
		fmt.Sprintf("master: %s", cfg.master),
		fmt.Sprintf("tag: %s", cfg.tag),
		fmt.Sprintf("test: %s", cfg.test),
		fmt.Sprintf("versionScheme: %s", nullOrValue(cfg.versionScheme)),
		fmt.Sprintf("versionType: %s", cfg.versionType),
	}

	content := strings.Join(lines, "\n") + "\n"

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		log.Fatalf(helper.FormatMessage("Error writing .gitcd config: %v", "error"), err)
	}
}
