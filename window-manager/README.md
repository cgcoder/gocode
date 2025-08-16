# macOS Window Manager

A lightweight Go application for managing window positions and sizes on macOS using global keyboard shortcuts.

## Features

- **Left Half**: Move focused window to left half of screen (`Ctrl+Cmd+Left`)
- **Right Half**: Move focused window to right half of screen (`Ctrl+Cmd+Right`) 
- **Fullscreen Toggle**: Toggle focused window between fullscreen and windowed mode (`Ctrl+Cmd+F`)

## Requirements

- macOS (tested on macOS 10.14+)
- Go 1.21 or later
- Xcode Command Line Tools (for CGO)
- Accessibility permissions

## Installation

1. **Clone the repository** (if not already done):
   ```bash
   git clone https://github.com/cgcoder/gocode.git
   cd gocode/window-manager
   ```

2. **Build the application**:
   ```bash
   go build -o window-manager .
   ```

## Setup

### Grant Accessibility Permissions

The application requires accessibility permissions to control other applications' windows.

1. Run the application once - it will prompt you about permissions
2. Go to **System Preferences** > **Security & Privacy** > **Privacy** > **Accessibility**
3. Click the lock icon and enter your password
4. Add the `window-manager` executable to the list of allowed applications
5. Restart the application

## Usage

### Hotkey Mode (Default)

1. **Start the window manager**:
   ```bash
   ./window-manager
   ```

2. **Use keyboard shortcuts**:
   - `Ctrl+Cmd+Left Arrow` - Move window to left half of screen
   - `Ctrl+Cmd+Right Arrow` - Move window to right half of screen  
   - `Ctrl+Cmd+F` - Toggle between fullscreen and windowed mode
   - `Ctrl+C` - Quit the application

### Interactive CLI Mode

For manual control or testing without hotkeys:

1. **Start in interactive mode**:
   ```bash
   ./window-manager --interactive
   ```

2. **Use CLI commands**:
   - `left` or `l` - Move window to left half
   - `right` or `r` - Move window to right half
   - `fullscreen` or `f` - Toggle fullscreen
   - `help` or `h` - Show available commands
   - `quit` or `q` - Exit the application

## How It Works

The application uses macOS's Accessibility API (AXUIElement) and Carbon Event Manager to:

1. **Window Management**: Access and manipulate window properties of the currently focused application
2. **Global Hotkeys**: Register system-wide keyboard shortcuts that work regardless of which application is active
3. **Screen Information**: Automatically detect screen dimensions and menu bar height for proper window positioning

## Technical Details

- Written in Go with CGO bindings to macOS frameworks
- Uses Cocoa, ApplicationServices, and Carbon frameworks
- Handles multiple monitor setups
- Respects macOS menu bar and dock positioning
- Graceful error handling for edge cases

## Troubleshooting

### "accessibility permissions required" Error
- Grant accessibility permissions as described in the Setup section
- Restart the application after granting permissions

### Hotkeys Not Working
- Ensure no other applications are using the same key combinations
- Check that the application is running in the foreground initially
- Verify accessibility permissions are granted

### Window Not Moving
- Some applications may restrict window manipulation
- Full-screen applications may not respond to positioning commands
- Try with standard applications like Safari, TextEdit, or Finder first

## Development

To modify or extend the functionality:

1. **Edit the Go code** for application logic
2. **Modify the C code** (in comments) for low-level macOS integration
3. **Rebuild** with `go build`

The C code uses:
- `AXUIElement` for window manipulation
- `EventHotKey` for global keyboard shortcuts
- `CGDisplay` functions for screen information

## License

This project follows the same license as the parent repository.