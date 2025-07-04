package helper

import (
	"context"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func Authenticate() (context.Context, github.Client) {
	config := ReadConfigFile()
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return ctx, *client
}
