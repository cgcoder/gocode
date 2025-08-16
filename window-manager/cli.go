package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CLI provides a command-line interface for manual window operations
type CLI struct {
	wm *WindowManager
}

// NewCLI creates a new CLI instance
func NewCLI(wm *WindowManager) *CLI {
	return &CLI{wm: wm}
}

// Start starts the interactive CLI
func (cli *CLI) Start() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nInteractive mode - Available commands:")
	cli.showHelp()

	for {
		fmt.Print("\nwm> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		command := strings.TrimSpace(strings.ToLower(input))

		switch command {
		case "left", "l":
			if err := cli.wm.MoveWindowToLeftHalf(); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Window moved to left half")
			}

		case "right", "r":
			if err := cli.wm.MoveWindowToRightHalf(); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Window moved to right half")
			}

		case "fullscreen", "full", "f":
			if err := cli.wm.ToggleWindowFullscreen(); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Window fullscreen toggled")
			}

		case "help", "h", "?":
			cli.showHelp()

		case "quit", "exit", "q":
			fmt.Println("Goodbye!")
			return

		case "":
			// Empty input, continue
			continue

		default:
			fmt.Printf("Unknown command: %s\n", command)
			fmt.Println("Type 'help' to see available commands")
		}
	}
}

// showHelp displays available CLI commands
func (cli *CLI) showHelp() {
	fmt.Println("  left, l      - Move focused window to left half")
	fmt.Println("  right, r     - Move focused window to right half")
	fmt.Println("  fullscreen, full, f - Toggle fullscreen")
	fmt.Println("  help, h, ?   - Show this help")
	fmt.Println("  quit, exit, q - Exit interactive mode")
}
