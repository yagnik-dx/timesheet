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
	shortPtr := flag.Bool("short", false, "Compact output format")

	flag.Usage = func() {
		fmt.Println("timesheet - Show git commit activity grouped by repository")
		fmt.Println("           Output is printed and copied to clipboard.\n")
		fmt.Println("Usage:")
		fmt.Println("  timesheet [flags]\n")
		fmt.Println("Flags:")
		fmt.Println("  --date YYYY-MM-DD     Show commits for specific date")
		fmt.Println("  --author EMAIL        Override detected git author")
		fmt.Println("  --path DIRECTORY      Scan repositories in directory")
		fmt.Println("  --short               Compact output format")
		fmt.Println("  --help                Show command help")
	}

	flag.Parse()

	if err := exec.Command("git", "--version").Run(); err != nil {
		fmt.Println("git command not found. Please install git.")
		os.Exit(1)
	}

	author := *authorPtr
	if author == "" {
		detectedAuthor, err := GetDefaultAuthor()
		if err != nil || detectedAuthor == "" {
			fmt.Println("Could not detect git author. Please specify --author.")
			os.Exit(1)
		}
		author = detectedAuthor
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

	report := GenerateReport(*datePtr, author, repos, *shortPtr)
	fmt.Println(report)

	if err := CopyToClipboard(report); err == nil {
		fmt.Println("Output copied to clipboard.")
	} else {
		fmt.Printf("Failed to copy to clipboard: %v\n", err)
	}
}
