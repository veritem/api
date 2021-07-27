package github

import (
	"context"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var Ctx = context.Background()

func Client() *githubv4.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")})
	tc := oauth2.NewClient(Ctx, ts)
	client := githubv4.NewClient(tc)
	return client
}
