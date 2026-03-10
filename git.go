package main

import (
	"os/exec"
	"strings"
)

// getAuthorFromGit runs `git config user.email` in the given directory (or global if dir is empty).
func getAuthorFromGit(dir string) (string, error) {
	args := []string{"config", "user.email"}
	if dir == "" {
		args = []string{"config", "--global", "user.email"}
	}
	cmd := exec.Command("git", args...)
	if dir != "" {
		cmd.Dir = dir
	}
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// GetDefaultAuthor gets the git author email: tries current dir, then global config, then from first repo if provided.
func GetDefaultAuthor(firstRepoPath string) (string, error) {
	// 1) Try current working directory (local config)
	if author, err := getAuthorFromGit("."); err == nil && author != "" {
		return author, nil
	}
	// 2) Try global config
	if author, err := getAuthorFromGit(""); err == nil && author != "" {
		return author, nil
	}
	// 3) Try from first discovered repo
	if firstRepoPath != "" {
		if author, err := getAuthorFromGit(firstRepoPath); err == nil && author != "" {
			return author, nil
		}
	}
	return "", nil
}

// GetCommits gets the commits for a specific author and date in a given repository.
func GetCommits(repoPath, date, author string) ([]string, error) {
	since := date + " 00:00:00"
	until := date + " 23:59:59"

	cmd := exec.Command("git", "log",
		"--since="+since,
		"--until="+until,
		"--author="+author,
		"--pretty=format:%s",
	)
	cmd.Dir = repoPath

	out, err := cmd.Output()
	if err != nil || len(out) == 0 {
		return nil, nil // Error might occur if the repository has no commits yet
	}

	outputStr := strings.TrimSpace(string(out))
	if outputStr == "" {
		return nil, nil
	}

	lines := strings.Split(outputStr, "\n")
	var result []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result, nil
}
