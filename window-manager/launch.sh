#!/bin/bash

# macOS Window Manager Launch Script

APP_NAME="window-manager"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
APP_PATH="$SCRIPT_DIR/$APP_NAME"

# Check if running on macOS
if [[ "$OSTYPE" != "darwin"* ]]; then
    echo "Error: This application only works on macOS"
    echo "Current OS: $OSTYPE"
    exit 1
fi

# Check if binary exists
if [ ! -f "$APP_PATH" ]; then
    echo "Building $APP_NAME..."
    cd "$SCRIPT_DIR"
    if ! go build -o "$APP_NAME" .; then
        echo "Build failed!"
        exit 1
    fi
    echo "Build complete"
fi

# Check for accessibility permissions
echo "Checking accessibility permissions..."
if ! "$APP_PATH" --interactive <<< "quit" 2>/dev/null >/dev/null; then
    echo ""
    echo "⚠️  Accessibility permissions required!"
    echo ""
    echo "Please grant accessibility permissions:"
    echo "1. Go to System Preferences > Security & Privacy > Privacy > Accessibility"
    echo "2. Click the lock icon and enter your password"
    echo "3. Add '$APP_PATH' to the allowed applications list"
    echo "4. Run this script again"
    echo ""
    echo "Would you like to open System Preferences now? (y/n)"
    read -r response
    if [[ "$response" == "y" || "$response" == "Y" ]]; then
        open "x-apple.systempreferences:com.apple.preference.security?Privacy_Accessibility"
    fi
    exit 1
fi

# Run the application
echo "Starting $APP_NAME..."
exec "$APP_PATH" "$@"