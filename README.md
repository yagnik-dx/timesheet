# timesheet

Show git commit activity grouped by repository. Output is printed and copied to the clipboard.

## Installation

### Windows

Pre-built `timesheet.exe` is provided. Use one of these:

**Option 1 — CMD or PowerShell (recommended)**  
Download and run the installer (it will download the exe and add it to your PATH):

```cmd
curl -fsSL -o install.bat https://raw.githubusercontent.com/yagnik-dx/timesheet/main/install.bat && install.bat
```

Or in PowerShell:

```powershell
Invoke-WebRequest -Uri https://raw.githubusercontent.com/yagnik-dx/timesheet/main/install.bat -OutFile install.bat -UseBasicParsing; .\install.bat
```

**Option 2 — Git Bash or WSL**

```bash
curl -sSL https://raw.githubusercontent.com/yagnik-dx/timesheet/main/install.sh | bash
```

**Option 3 — Manual**  
Download [timesheet.exe](https://github.com/yagnik-dx/timesheet/blob/main/timesheet.exe) and put it in a folder on your PATH (e.g. `%USERPROFILE%\bin`).

After installing, close and reopen your terminal, then run `timesheet --help`.

### Linux & macOS (build from source)

Pre-built binaries are not provided. You need Go installed:

```bash
git clone https://github.com/yagnik-dx/timesheet.git
cd timesheet
go build -o timesheet .
sudo mv timesheet /usr/local/bin/
```

Or install via Go:

```bash
git clone https://github.com/yagnik-dx/timesheet.git && cd timesheet && go install .
```

(Make sure `$GOPATH/bin` or `$GOBIN` is in your PATH.)

## Requirements

- **Windows:** [Git](https://git-scm.com/) (must be on your PATH for the tool to work).
- **Linux & macOS:** [Go](https://go.dev/dl/) 1.x and [Git](https://git-scm.com/).

## Usage

```bash
timesheet [flags]
```

| Flag | Description |
|------|-------------|
| `--date YYYY-MM-DD` | Show commits for a specific date (default: today) |
| `--author EMAIL` | Override detected git author |
| `--path DIRECTORY` | Scan repositories in this directory (default: current directory) |
| `--short` | Compact output format |
| `--help` | Show command help |

## Examples

```bash
timesheet
timesheet --date 2026-03-10
timesheet --path C:\Projects
timesheet --short
```

## Updating

- **Windows:** Run the install script again, or replace `timesheet.exe` in your install directory.
- **Linux & macOS:** Pull the latest changes and rebuild (`go build` or `go install`).
