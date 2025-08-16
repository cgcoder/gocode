//go:build darwin
// +build darwin

package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework ApplicationServices -framework Carbon

#import <Cocoa/Cocoa.h>
#import <ApplicationServices/ApplicationServices.h>
#import <Carbon/Carbon.h>

// Global variables for window management
CGSize screenSize;
CGFloat menuBarHeight = 24; // Standard macOS menu bar height

// Initialize screen dimensions
void initScreenInfo() {
    CGDirectDisplayID mainDisplay = CGMainDisplayID();
    screenSize = CGDisplayPixelsHigh(mainDisplay) > 0 ?
        CGSizeMake(CGDisplayPixelsWide(mainDisplay), CGDisplayPixelsHigh(mainDisplay)) :
        CGSizeMake(1920, 1080); // fallback

    // Get actual menu bar height
    NSScreen *mainScreen = [NSScreen mainScreen];
    if (mainScreen) {
        CGFloat screenHeight = mainScreen.frame.size.height;
        CGFloat visibleHeight = mainScreen.visibleFrame.size.height;
        menuBarHeight = screenHeight - visibleHeight - mainScreen.visibleFrame.origin.y;
    }
}

// Get the frontmost window
AXUIElementRef getFrontmostWindow() {
    // Get frontmost application
    NSWorkspace *workspace = [NSWorkspace sharedWorkspace];
    NSArray *runningApps = [workspace runningApplications];

    for (NSRunningApplication *app in runningApps) {
        if ([app isActive]) {
            pid_t pid = [app processIdentifier];
            AXUIElementRef appRef = AXUIElementCreateApplication(pid);

            if (appRef) {
                CFTypeRef windowRef;
                AXError result = AXUIElementCopyAttributeValue(appRef, kAXFocusedWindowAttribute, &windowRef);
                CFRelease(appRef);

                if (result == kAXErrorSuccess && windowRef) {
                    return (AXUIElementRef)windowRef;
                }
            }
            break;
        }
    }
    return NULL;
}

// Set window position and size
int setWindowFrame(AXUIElementRef windowRef, CGFloat x, CGFloat y, CGFloat width, CGFloat height) {
    if (!windowRef) return 0;

    // Create position point
    CGPoint position = CGPointMake(x, y);
    AXValueRef positionValue = AXValueCreate(kAXValueCGPointType, &position);

    // Create size
    CGSize size = CGSizeMake(width, height);
    AXValueRef sizeValue = AXValueCreate(kAXValueCGSizeType, &size);

    // Set position and size
    AXError posResult = AXUIElementSetAttributeValue(windowRef, kAXPositionAttribute, positionValue);
    AXError sizeResult = AXUIElementSetAttributeValue(windowRef, kAXSizeAttribute, sizeValue);

    CFRelease(positionValue);
    CFRelease(sizeValue);

    return (posResult == kAXErrorSuccess && sizeResult == kAXErrorSuccess) ? 1 : 0;
}

// Move window to left half of screen
int moveWindowToLeftHalf() {
    initScreenInfo();
    AXUIElementRef window = getFrontmostWindow();
    if (!window) return 0;

    CGFloat width = screenSize.width / 2;
    CGFloat height = screenSize.height - menuBarHeight;
    int result = setWindowFrame(window, 0, menuBarHeight, width, height);

    CFRelease(window);
    return result;
}

// Move window to right half of screen
int moveWindowToRightHalf() {
    initScreenInfo();
    AXUIElementRef window = getFrontmostWindow();
    if (!window) return 0;

    CGFloat width = screenSize.width / 2;
    CGFloat height = screenSize.height - menuBarHeight;
    CGFloat x = screenSize.width / 2;
    int result = setWindowFrame(window, x, menuBarHeight, width, height);

    CFRelease(window);
    return result;
}

// Toggle window fullscreen
int toggleWindowFullscreen() {
    initScreenInfo();
    AXUIElementRef window = getFrontmostWindow();
    if (!window) return 0;

    // Try to get current size to determine if already fullscreen
    AXValueRef sizeValue;
    CGSize currentSize;
    AXError result = AXUIElementCopyAttributeValue(window, kAXSizeAttribute, (CFTypeRef*)&sizeValue);

    if (result == kAXErrorSuccess && sizeValue) {
        AXValueGetValue(sizeValue, kAXValueCGSizeType, &currentSize);
        CFRelease(sizeValue);

        // Check if already fullscreen (or close to it)
        if (currentSize.width >= screenSize.width - 10 && currentSize.height >= screenSize.height - menuBarHeight - 10) {
            // Currently fullscreen, restore to center with reasonable size
            CGFloat width = screenSize.width * 0.8;
            CGFloat height = (screenSize.height - menuBarHeight) * 0.8;
            CGFloat x = (screenSize.width - width) / 2;
            CGFloat y = menuBarHeight + (screenSize.height - menuBarHeight - height) / 2;
            result = setWindowFrame(window, x, y, width, height);
        } else {
            // Not fullscreen, make it fullscreen
            result = setWindowFrame(window, 0, menuBarHeight, screenSize.width, screenSize.height - menuBarHeight);
        }
    }

    CFRelease(window);
    return result ? 1 : 0;
}

// Global hotkey handling
EventHotKeyRef leftHalfHotKey;
EventHotKeyRef rightHalfHotKey;
EventHotKeyRef fullscreenHotKey;

// Hotkey handler
OSStatus hotKeyHandler(EventHandlerCallRef nextHandler, EventRef theEvent, void *userData) {
    EventHotKeyID hotKeyID;
    GetEventParameter(theEvent, kEventParamDirectObject, typeEventHotKeyID, NULL, sizeof(hotKeyID), NULL, &hotKeyID);

    switch (hotKeyID.id) {
        case 1: // Left half
            moveWindowToLeftHalf();
            break;
        case 2: // Right half
            moveWindowToRightHalf();
            break;
        case 3: // Fullscreen
            toggleWindowFullscreen();
            break;
    }

    return noErr;
}

// Register global hotkeys
int registerHotkeys() {
    EventTypeSpec eventType;
    eventType.eventClass = kEventClassKeyboard;
    eventType.eventKind = kEventHotKeyPressed;

    InstallApplicationEventHandler(&hotKeyHandler, 1, &eventType, NULL, NULL);

    // Register Ctrl+Cmd+Left (left half)
    EventHotKeyID leftHalfID = {1, 1};
    RegisterEventHotKey(kVK_LeftArrow, controlKey | cmdKey, leftHalfID, GetApplicationEventTarget(), 0, &leftHalfHotKey);

    // Register Ctrl+Cmd+Right (right half)
    EventHotKeyID rightHalfID = {2, 2};
    RegisterEventHotKey(kVK_RightArrow, controlKey | cmdKey, rightHalfID, GetApplicationEventTarget(), 0, &rightHalfHotKey);

    // Register Ctrl+Cmd+F (fullscreen)
    EventHotKeyID fullscreenID = {3, 3};
    RegisterEventHotKey(kVK_F, controlKey | cmdKey, fullscreenID, GetApplicationEventTarget(), 0, &fullscreenHotKey);

    return 1;
}

// Unregister hotkeys
void unregisterHotkeys() {
    if (leftHalfHotKey) UnregisterEventHotKey(leftHalfHotKey);
    if (rightHalfHotKey) UnregisterEventHotKey(rightHalfHotKey);
    if (fullscreenHotKey) UnregisterEventHotKey(fullscreenHotKey);
}

// Check accessibility permissions
int checkAccessibilityPermissions() {
    return AXIsProcessTrusted() ? 1 : 0;
}
*/
import "C"
import (
	"errors"
	"fmt"
	"runtime"
)

// WindowManager handles macOS window management operations
type WindowManager struct {
	initialized bool
}

// NewWindowManager creates a new window manager instance
func NewWindowManager() (*WindowManager, error) {
	if runtime.GOOS != "darwin" {
		return nil, errors.New("this application only works on macOS")
	}

	// Check accessibility permissions
	if C.checkAccessibilityPermissions() == 0 {
		return nil, errors.New("accessibility permissions required. Please grant accessibility permissions to this app in System Preferences > Security & Privacy > Privacy > Accessibility")
	}

	wm := &WindowManager{
		initialized: true,
	}

	return wm, nil
}

// RegisterShortcuts registers global keyboard shortcuts
func (wm *WindowManager) RegisterShortcuts() error {
	if !wm.initialized {
		return errors.New("window manager not initialized")
	}

	result := C.registerHotkeys()
	if result == 0 {
		return errors.New("failed to register hotkeys")
	}

	return nil
}

// MoveWindowToLeftHalf moves the focused window to the left half of the screen
func (wm *WindowManager) MoveWindowToLeftHalf() error {
	if !wm.initialized {
		return errors.New("window manager not initialized")
	}

	result := C.moveWindowToLeftHalf()
	if result == 0 {
		return errors.New("failed to move window to left half")
	}

	return nil
}

// MoveWindowToRightHalf moves the focused window to the right half of the screen
func (wm *WindowManager) MoveWindowToRightHalf() error {
	if !wm.initialized {
		return errors.New("window manager not initialized")
	}

	result := C.moveWindowToRightHalf()
	if result == 0 {
		return errors.New("failed to move window to right half")
	}

	return nil
}

// ToggleWindowFullscreen toggles the focused window between fullscreen and windowed mode
func (wm *WindowManager) ToggleWindowFullscreen() error {
	if !wm.initialized {
		return errors.New("window manager not initialized")
	}

	result := C.toggleWindowFullscreen()
	if result == 0 {
		return errors.New("failed to toggle window fullscreen")
	}

	return nil
}

// Close cleans up the window manager
func (wm *WindowManager) Close() {
	if wm.initialized {
		C.unregisterHotkeys()
		wm.initialized = false
	}
}

// CheckAccessibilityPermissions checks if the app has accessibility permissions
func CheckAccessibilityPermissions() bool {
	return C.checkAccessibilityPermissions() == 1
}
