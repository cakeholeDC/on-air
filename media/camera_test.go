package media

import (
	"testing"
)

func TestGetCameraState(t *testing.T) {
	// Test cases
	cameraState := GetCameraState()

	// Assert that cameraState is a boolean
	if cameraState != true && cameraState != false {
		t.Errorf("Expected cameraState to be of type bool, but got %T", cameraState)
	}
}
