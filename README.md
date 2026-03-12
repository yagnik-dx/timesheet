# timesheet

**timesheet** is a lightweight Go CLI that scans Git repositories and lists commit messages grouped by repository — ideal for daily standup reports.

Output is automatically copied to your clipboard on every run.

---

## Requirements

- Git on PATH
- Windows (pre-built binary included)

---

## Setup

No Go required — a pre-built Windows binary is included.

**CMD**
```cmd
git clone https://github.com/yagnik-dx/timesheet.git
cd timesheet
install.bat
```

**PowerShell**
```powershell
git clone https://github.com/yagnik-dx/timesheet.git
cd timesheet
.\install.bat
```

This copies `bin\timesheet.exe` to `%USERPROFILE%\bin` and adds it to your PATH.

> Open a new terminal after installation and run `timesheet` to verify.

---

## Usage

```
timesheet [flags]
```

| Flag | Description |
|------|-------------|
| `--date YYYY-MM-DD` | Show commits for a specific date (default: today) |
| `--author EMAIL` | Override detected git author |
| `--path DIRECTORY` | Scan repositories in this directory (default: current directory) |
| `--copy=false` | Disable clipboard copy |
| `--short` | Compact output format |
| `--help` | Show command help |

---

## Examples

```bash
timesheet
timesheet --date 2026-03-10
timesheet --author dev@company.com
timesheet --path C:\Projects
timesheet --short
timesheet --copy=false
```

---

## Output

Default:
```
Date: 2026-03-12
Author: you@example.com

auth-service
- fix login validation
- add rate limiting

billing-service
- optimize invoice query
```

With `--short`:
```
auth-service: fix login validation
auth-service: add rate limiting
billing-service: optimize invoice query
```
