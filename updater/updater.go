package updater

import (
	"context"
	"fmt"
	"os"

	"github.com/creativeprojects/go-selfupdate"
)

const repoSlug = "Arondy/go_updates"

func Update(version string) error {
	ctx := context.Background()
	slug := selfupdate.ParseSlug(repoSlug)

	fmt.Println("Checking for updates...")

	source, err := selfupdate.NewGitHubSource(selfupdate.GitHubConfig{
		APIToken: os.Getenv("GITHUB_TOKEN"),
	})
	if err != nil {
		return fmt.Errorf("failed to create GitHub source: %w", err)
	}

	updater, err := selfupdate.NewUpdater(selfupdate.Config{
		Source:    source,
		Validator: &selfupdate.ChecksumValidator{UniqueFilename: "checksums.txt"},
	})
	if err != nil {
		return fmt.Errorf("updater init: %w", err)
	}

	latest, found, err := updater.DetectLatest(ctx, slug)
	if err != nil {
		return fmt.Errorf("failed to detect latest version: %w", err)
	}
	if !found {
		return fmt.Errorf("no releases found")
	}

	if latest.LessOrEqual(version) {
		fmt.Println("Already latest:", version)
		return nil
	}

	fmt.Printf("Found new version: %s (current: %s)\n", latest.Version(), version)
	fmt.Println("Updating...")

	release, err := updater.UpdateSelf(ctx, version, slug)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	fmt.Println("Updated successfully!")
	fmt.Println("Notes:\n", release.ReleaseNotes)
	return nil
}
