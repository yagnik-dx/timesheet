package main

import (
	"os"
	"path/filepath"
)

// FindRepositories scans the given root path recursively and returns Git repository paths.
func FindRepositories(root string) ([]string, error) {
	var repos []string

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // ignore errors
		}

		if !d.IsDir() {
			return nil
		}

		gitPath := filepath.Join(path, ".git")
		info, err := os.Stat(gitPath)
		if err == nil && info.IsDir() {
			repos = append(repos, path)
			return filepath.SkipDir // skip scanning the contents of this repository
		}

		return nil
	})

	return repos, err
}
