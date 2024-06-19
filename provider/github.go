package provider

import (
	"context"

	"github.com/google/go-github/v62/github"
)

func GetLatestRelease() (*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "biel-correa", "correa-cli")
	return release, err
}
