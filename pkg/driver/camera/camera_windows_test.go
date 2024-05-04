//go:build windows
// +build windows

package camera

import (
	"testing"
)

func TestCameraInitialization(t *testing.T) {
	// Call the Initialize function and check for errors
	Initialize()
	// Add assertions or checks here to ensure that the initialization was successful
}
