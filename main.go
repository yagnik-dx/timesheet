package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	datePtr := flag.String("date", time.Now().Format("2006-01-02"), "Show commits for specific date")
	authorPtr := flag.String("author", "", "Override detected git author")
	pathPtr := flag.String("path", ".", "Scan repositories in directory")
	copyPtr := flag.Bool("copy", true, "Copy output to clipboard")
	shortPtr := flag.Bool("short", false, "Compact output format")
	helpPtr := flag.Bool("help", false, "Show command help")

	flag.Usage = func() {
		fmt.Println("timesheet - Show git commit activity grouped by repository")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  timesheet [flags]")
		fmt.Println()
		fmt.Println("Flags:")
		fmt.Println("  --date YYYY-MM-DD     Show commits for specific date")
		fmt.Println("  --author EMAIL        Override detected git author")
		fmt.Println("  --path DIRECTORY      Scan repositories in directory")
		fmt.Println("  --copy                Copy output to clipboard")
		fmt.Println("  --short               Compact output format")
		fmt.Println("  --help                Show command help")
	}

	flag.Parse()

	if *helpPtr {
		flag.Usage()
		return
	}

	if err := exec.Command("git", "--version").Run(); err != nil {
		fmt.Println("git command not found. Please install git.")
		os.Exit(1)
	}

	repos, err := FindRepositories(*pathPtr)
	if err != nil {
		fmt.Printf("Error scanning directories: %v\n", err)
		os.Exit(1)
	}

	if len(repos) == 0 {
		fmt.Println("No git repositories found.")
		return
	}

	var firstRepo string
	if len(repos) > 0 {
		firstRepo = repos[0]
	}
	author := *authorPtr
	if author == "" {
		detectedAuthor, err := GetDefaultAuthor(firstRepo)
		if err != nil || detectedAuthor == "" {
			fmt.Println("Could not detect git author. Please specify --author or set git config user.email.")
			os.Exit(1)
		}
		author = detectedAuthor
	}

	report := GenerateReport(*datePtr, author, repos, *shortPtr)
	fmt.Println(report)

	if *copyPtr {
		if err := CopyToClipboard(report); err == nil {
			fmt.Println("Output copied to clipboard.")
		} else {
			fmt.Printf("Failed to copy to clipboard: %v\n", err)
		}
	}
}
