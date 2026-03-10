# timesheet

Show git commit activity grouped by repository. Output is printed and copied to the clipboard.

## Installation

### Windows (one-line install)

Pre-built `timesheet.exe` is provided. In Git Bash or WSL:

```bash
curl -sSL https://raw.githubusercontent.com/yagnik-dx/timesheet/main/install.sh | bash
```

Then run:

```bash
timesheet --help
```

**Manual:** Download [timesheet.exe](https://github.com/yagnik-dx/timesheet/blob/main/bin/timesheet.exe) from the `bin/` folder and place it in a directory on your PATH.

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
