package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	webhookURL = "https://webhook-test.com/ec1852b6a8ccd5cdd91e2bb8dc162c0e"
	tempFile   = "/tmp/neovim_buffer.txt"
)

func main() {
	// Create a temporary file
	if err := os.WriteFile(tempFile, []byte{}, 0644); err != nil {
		fmt.Printf("Error creating temporary file: %v\n", err)
		return
	}
	defer os.Remove(tempFile)

	// Get the initial file info
	initialInfo, err := os.Stat(tempFile)
	if err != nil {
		fmt.Printf("Error getting initial file info: %v\n", err)
		return
	}

	// Open the file in Neovim in insert mode
	cmd := exec.Command("nvim", "+startinsert", tempFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running Neovim: %v\n", err)
		return
	}

	// Get the file info after editing
	afterInfo, err := os.Stat(tempFile)
	if err != nil {
		fmt.Printf("Error getting file info after editing: %v\n", err)
		return
	}

	// Check if the file was modified
	if afterInfo.ModTime() == initialInfo.ModTime() || afterInfo.Size() == 0 {
		fmt.Println("File was not modified or is empty. Not sending request.")
		return
	}

	// Read the content of the file
	content, err := os.ReadFile(tempFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Trim leading and trailing whitespace
	trimmedContent := strings.TrimSpace(string(content))

	// Check if the trimmed content is empty
	if trimmedContent == "" {
		fmt.Println("File content is empty after trimming. Not sending request.")
		return
	}

	// Send the content to the webhook
	if err := sendToWebhook([]byte(trimmedContent)); err != nil {
		fmt.Printf("Error sending to webhook: %v\n", err)
		return
	}

	fmt.Println("Content sent to webhook successfully!")
}

func sendToWebhook(content []byte) error {
	resp, err := http.Post(webhookURL, "text/plain", bytes.NewReader(content))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
