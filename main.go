package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

const (
	webhookURL = "https://webhook-test.com/ec1852b6a8ccd5cdd91e2bb8dc162c0e"
	tempFile   = "/tmp/neovim_buffer.txt"
)

func main() {
	// Create a temporary file
	if err := ioutil.WriteFile(tempFile, []byte{}, 0644); err != nil {
		fmt.Printf("Error creating temporary file: %v\n", err)
		return
	}
	defer os.Remove(tempFile)

	// Open the file in Neovim
	cmd := exec.Command("nvim", tempFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running Neovim: %v\n", err)
		return
	}

	// Read the content of the file
	content, err := ioutil.ReadFile(tempFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Send the content to the webhook
	if err := sendToWebhook(content); err != nil {
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
