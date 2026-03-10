package main

import (
	"os/exec"
	"strings"
)

// GetDefaultAuthor gets the default git author email.
func GetDefaultAuthor() (string, error) {
	cmd := exec.Command("git", "config", "user.email")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
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
