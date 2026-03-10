package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type RepoCommits struct {
	Name    string
	Commits []string
}

// formatDisplayDate converts YYYY-MM-DD to "06 March 2026" format.
func formatDisplayDate(date string) string {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return date
	}
	return t.Format("02 January 2006")
}

// GenerateReport formats the commit messages grouped by repository.
func GenerateReport(date, author string, repos []string, short bool) string {
	var repoCommitsList []RepoCommits
	for _, repoPath := range repos {
		commits, _ := GetCommits(repoPath, date, author)
		if len(commits) > 0 {
			repoCommitsList = append(repoCommitsList, RepoCommits{
				Name:    filepath.Base(repoPath),
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
		sb.WriteString("------------------\n")
		sb.WriteString(formatDisplayDate(date) + "\n")
		sb.WriteString("------------------\n")
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
