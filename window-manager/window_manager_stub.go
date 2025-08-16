//go:build !darwin
// +build !darwin

package main

import (
	"errors"
	"runtime"
)

// WindowManager handles window management operations (stub for non-macOS)
type WindowManager struct {
	initialized bool
}

// NewWindowManager creates a new window manager instance
func NewWindowManager() (*WindowManager, error) {
	return nil, errors.New("this application only works on macOS, current OS: " + runtime.GOOS)
}

// RegisterShortcuts registers global keyboard shortcuts (stub)
func (wm *WindowManager) RegisterShortcuts() error {
	return errors.New("not supported on " + runtime.GOOS)
}

// MoveWindowToLeftHalf moves the focused window to the left half of the screen (stub)
func (wm *WindowManager) MoveWindowToLeftHalf() error {
	return errors.New("not supported on " + runtime.GOOS)
}

// MoveWindowToRightHalf moves the focused window to the right half of the screen (stub)
func (wm *WindowManager) MoveWindowToRightHalf() error {
	return errors.New("not supported on " + runtime.GOOS)
}

// ToggleWindowFullscreen toggles the focused window between fullscreen and windowed mode (stub)
func (wm *WindowManager) ToggleWindowFullscreen() error {
	return errors.New("not supported on " + runtime.GOOS)
}

// Close cleans up the window manager (stub)
func (wm *WindowManager) Close() {
	// No-op for non-macOS
}

// CheckAccessibilityPermissions checks if the app has accessibility permissions (stub)
func CheckAccessibilityPermissions() bool {
	return false
}
