package helper

import (
	"context"
	"log"

	"github.com/google/go-github/v73/github"
	"golang.org/x/oauth2"
)

func Authenticate() (context.Context, github.Client) {
	config, err := ReadConfigFile()
	if err != nil {
		log.Fatalf(FormatMessage("Authenticate: Error opening config file: %v\n", "error"), err)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return ctx, *client
}
