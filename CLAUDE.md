# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Run

```bash
go build -o timesheet .
./timesheet --help
```

Cross-platform builds:
```bash
GOOS=windows GOARCH=amd64 go build -o bin/timesheet.exe .
GOOS=linux   GOARCH=amd64 go build -o bin/timesheet-linux .
GOOS=darwin  GOARCH=amd64 go build -o bin/timesheet-macos .
GOOS=darwin  GOARCH=arm64 go build -o bin/timesheet-macos-arm64 .
```

No external dependencies — stdlib only (`go.mod` declares `module timesheet`).

## Architecture

Single `main` package with no third-party imports. Each file has a clear responsibility:

- **main.go** — CLI flag parsing, orchestration: scans repos → detects author → generates report → optionally copies to clipboard
- **scanner.go** — `FindRepositories(root)`: walks directory tree, finds `.git` dirs, skips repo internals
- **git.go** — `GetDefaultAuthor(firstRepo)`: tries local → global → first-repo git config; `GetCommits(repo, date, author)`: runs `git log` with `--since`/`--until`/`--author` filtering
- **report.go** — `GenerateReport(date, author, repos, short)`: calls `GetCommits` per repo, sorts results alphabetically, formats as full (Date/Author header + bullet lists) or compact (`repo: commit` lines)
- **clipboard.go** — `CopyToClipboard(text)`: dispatches to `clip` (Windows), `pbcopy` (macOS), or `xclip`/`xsel` (Linux)

## Data Flow

`FindRepositories` → `GetDefaultAuthor` → `GenerateReport` (calls `GetCommits` per repo) → optional `CopyToClipboard`

All git operations shell out to the `git` binary; there is no libgit2 or go-git dependency.
