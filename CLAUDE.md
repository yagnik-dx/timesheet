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

## Lint & Format

```bash
go vet ./...
gofmt -w .
```

No tests exist in this codebase.

## Architecture

Single `main` package, no third-party imports. Each file has a clear responsibility:

- **main.go** — CLI flag parsing, orchestration: scans repos → detects author → generates report → optionally copies to clipboard
- **scanner.go** — `FindRepositories(root)`: walks directory tree, finds `.git` dirs; calls `filepath.SkipDir` after finding one so nested repos are not traversed
- **git.go** — `GetDefaultAuthor(firstRepo)`: tries local → global → first-repo git config; `GetCommits(repo, date, author)`: runs `git log --pretty=format:%s` filtered by `--since`/`--until`/`--author`; returns `nil, nil` on error (treated as no commits)
- **report.go** — `GenerateReport(date, author, repos, short)`: calls `GetCommits` per repo, sorts alphabetically by `filepath.Base` of the repo path, formats as full (Date/Author header + bullet lists) or compact (`repo: commit` lines)
- **clipboard.go** — `CopyToClipboard(text)`: dispatches to `clip` (Windows), `pbcopy` (macOS), or `xclip`/`xsel` (Linux); Linux requires one of these to be installed

## Key Behavioral Notes

- `--copy` defaults to `true`; disable with `--copy=false`
- Author is matched via `--author` substring (passed directly to `git log --author=`), so partial email matches work
- `WalkDir` errors are silently ignored; inaccessible directories are skipped without failing
- The Windows installer (`install.bat`) copies `bin\timesheet.exe` to `%USERPROFILE%\bin` and patches the user PATH via PowerShell
