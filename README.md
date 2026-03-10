# timesheet

**timesheet** is a lightweight Go CLI that scans Git repositories and lists commit messages grouped by repository—ideal for daily standup reports.

---

## Features

- Scan directories recursively to discover Git repos
- Auto-detect Git user email (local → global → first repo)
- Show commit messages grouped by repository
- Default: **today's commits**
- Filter by **date**, **author**, and **path**
- Optional copy to clipboard (`--copy`)
- Compact output with `--short`

---

## Installation

Run:

```bash
curl -sSL https://raw.githubusercontent.com/yagnik-dx/timesheet/main/install.sh | bash
```

Then:

```bash
timesheet --help
```

---

## Usage

```bash
timesheet [flags]
```

| Flag | Description |
|------|-------------|
| `--date YYYY-MM-DD` | Show commits for a specific date (default: today) |
| `--author EMAIL` | Override detected git author |
| `--path DIRECTORY` | Scan repositories in this directory (default: current directory) |
| `--copy` | Copy output to clipboard |
| `--short` | Compact output format |
| `--help` | Show command help |

---

## Examples

```bash
timesheet
timesheet --date 2026-03-10
timesheet --author dev@company.com
timesheet --path C:\Projects
timesheet --copy
timesheet --short
```

---

## Build manually

If you have Go installed and want to build locally:

```bash
go build -o timesheet .
./timesheet --help
```

To build for multiple platforms:

```bash
GOOS=windows GOARCH=amd64 go build -o bin/timesheet.exe .
GOOS=linux   GOARCH=amd64 go build -o bin/timesheet-linux .
GOOS=darwin  GOARCH=amd64 go build -o bin/timesheet-macos .
GOOS=darwin  GOARCH=arm64 go build -o bin/timesheet-macos-arm64 .
```

---

## Requirements

- **Git** on PATH. For `--copy`: Windows `clip`, macOS `pbcopy`, or Linux `xclip`/`xsel`.
