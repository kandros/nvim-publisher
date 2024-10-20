# Neovim Webhook CLI

This is a simple Golang CLI application that opens a buffer in Neovim and sends the content to a webhook when the file is saved and closed.

## Features

- Opens a temporary buffer in Neovim in insert mode
- Sends the buffer content to a specified webhook URL after closing Neovim
- Removes leading and trailing whitespace from the content before sending
- Does not send a request if the file is empty or not modified

## Usage

1. Ensure you have Go and Neovim installed on your system.
2. Build the application:   ```
   go build -o neovim-webhook   ```
3. Run the application:   ```
   ./neovim-webhook   ```
4. Edit the content in Neovim (you'll start in insert mode), save (`:w`), and quit (`:q`).
5. If the file is not empty and was modified, the application will send the content to the specified webhook URL.

## Configuration

The webhook URL is currently hardcoded in the `main.go` file. To change it, modify the `webhookURL` constant.

## Note

This application creates a temporary file in the `/tmp` directory. The file is removed after the application finishes executing.
