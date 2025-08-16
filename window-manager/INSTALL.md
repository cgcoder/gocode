# Quick Installation Guide

## For macOS Users

1. **Navigate to the window-manager directory**:
   ```bash
   cd window-manager
   ```

2. **Use the launch script (recommended)**:
   ```bash
   ./launch.sh
   ```
   
   The launch script will:
   - Build the application if needed
   - Check for required permissions
   - Guide you through setup if needed
   - Launch the application

3. **Or build and run manually**:
   ```bash
   make build
   ./window-manager
   ```

## For Development

1. **Run tests**:
   ```bash
   go test -v
   ```

2. **Build**:
   ```bash
   make build
   ```

3. **Clean**:
   ```bash
   make clean
   ```

## Modes

- **Default**: Global hotkeys (Ctrl+Cmd+Arrow keys)
- **Interactive**: Manual commands (`./window-manager --interactive`)

See README.md for complete documentation.