package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func repoDisplayName(repoPath string) string {
	abs, err := filepath.Abs(repoPath)
	if err != nil {
		return filepath.Base(repoPath)
	}
	return filepath.Base(abs)
}

type RepoCommits struct {
	Name    string
	Commits []string
}

// formatDisplayDate returns the date in YYYY-MM-DD for header (spec format).
func formatDisplayDate(date string) string {
	if _, err := time.Parse("2006-01-02", date); err != nil {
		return date
	}
	return date
}

// GenerateReport formats the commit messages grouped by repository.
func GenerateReport(date, author string, repos []string, short bool) string {
	var repoCommitsList []RepoCommits
	for _, repoPath := range repos {
		commits, _ := GetCommits(repoPath, date, author)
		if len(commits) > 0 {
			repoCommitsList = append(repoCommitsList, RepoCommits{
				Name:    repoDisplayName(repoPath),
				Commits: commits,
			})
		}
	}

	if len(repoCommitsList) == 0 {
		return fmt.Sprintf("No commits found for %s", formatDisplayDate(date))
	}

	// Sort by repository name alphabetically
	sort.Slice(repoCommitsList, func(i, j int) bool {
		return repoCommitsList[i].Name < repoCommitsList[j].Name
	})

	var sb strings.Builder

	if short {
		for i, rc := range repoCommitsList {
			for j, commit := range rc.Commits {
				sb.WriteString(fmt.Sprintf("%s: %s", rc.Name, commit))
				if i < len(repoCommitsList)-1 || j < len(rc.Commits)-1 {
					sb.WriteString("\n")
				}
			}
		}
	} else {
		sb.WriteString(fmt.Sprintf("Date: %s\n", formatDisplayDate(date)))
		sb.WriteString(fmt.Sprintf("Author: %s\n\n", author))

		for i, rc := range repoCommitsList {
			sb.WriteString(rc.Name + "\n")
			for _, commit := range rc.Commits {
				sb.WriteString(fmt.Sprintf("- %s\n", commit))
			}
			if i < len(repoCommitsList)-1 {
				sb.WriteString("\n")
			}
		}
	}

	return strings.TrimRight(sb.String(), "\n")
}
