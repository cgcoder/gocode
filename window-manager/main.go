package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Parse command line flags
	interactive := flag.Bool("interactive", false, "Run in interactive CLI mode")
	flag.Parse()

	fmt.Println("macOS Window Manager")
	fmt.Println("====================")

	// Initialize window manager
	wm, err := NewWindowManager()
	if err != nil {
		log.Fatalf("Failed to initialize window manager: %v", err)
	}
	defer wm.Close()

	if *interactive {
		// Run in interactive CLI mode
		fmt.Println("Running in interactive mode...")
		fmt.Println("Use commands to manually control windows")

		cli := NewCLI(wm)
		cli.Start()
	} else {
		// Run in hotkey mode (default)
		fmt.Println("Keyboard shortcuts:")
		fmt.Println("  Ctrl+Cmd+Left  - Move window to left half")
		fmt.Println("  Ctrl+Cmd+Right - Move window to right half")
		fmt.Println("  Ctrl+Cmd+F     - Toggle fullscreen")
		fmt.Println("  Ctrl+C         - Quit")
		fmt.Println()
		fmt.Println("Use --interactive flag to run in CLI mode")

		// Register keyboard shortcuts
		if err := wm.RegisterShortcuts(); err != nil {
			log.Fatalf("Failed to register shortcuts: %v", err)
		}

		fmt.Println("Window manager started. Press Ctrl+C to quit.")

		// Wait for interrupt signal
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c

		fmt.Println("\nShutting down window manager...")
	}
}
