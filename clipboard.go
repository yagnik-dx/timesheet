package main

import (
	"os/exec"
	"strings"
)

// CopyToClipboard copies text to the system clipboard using the clip command on Windows.
func CopyToClipboard(text string) error {
	cmd := exec.Command("clip")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}
