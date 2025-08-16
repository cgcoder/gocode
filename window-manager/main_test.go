package main

import (
	"runtime"
	"testing"
)

func TestWindowManagerCreation(t *testing.T) {
	wm, err := NewWindowManager()

	if runtime.GOOS == "darwin" {
		// On macOS, should work (assuming permissions)
		if err != nil && err.Error() != "accessibility permissions required. Please grant accessibility permissions to this app in System Preferences > Security & Privacy > Privacy > Accessibility" {
			t.Fatalf("Unexpected error on macOS: %v", err)
		}
		if wm != nil {
			defer wm.Close()
		}
	} else {
		// On non-macOS, should fail with specific error
		if err == nil {
			t.Fatal("Expected error on non-macOS platform")
		}
		if wm != nil {
			t.Fatal("WindowManager should be nil on non-macOS platform")
		}
		expected := "this application only works on macOS, current OS: " + runtime.GOOS
		if err.Error() != expected {
			t.Fatalf("Expected error '%s', got '%s'", expected, err.Error())
		}
	}
}

func TestCLICreation(t *testing.T) {
	// CLI should be creatable even with nil WindowManager
	cli := NewCLI(nil)
	if cli == nil {
		t.Fatal("CLI should not be nil")
	}
	if cli.wm != nil {
		t.Fatal("CLI WindowManager should be nil when passed nil")
	}
}

func TestCheckAccessibilityPermissions(t *testing.T) {
	// Should not panic on any platform
	result := CheckAccessibilityPermissions()

	if runtime.GOOS != "darwin" {
		// On non-macOS, should always return false
		if result {
			t.Fatal("CheckAccessibilityPermissions should return false on non-macOS")
		}
	}
	// On macOS, result depends on actual system permissions, so we don't test the value
}
