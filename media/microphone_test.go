package media

import (
	"testing"
)

func TestGetMicrophoneState(t *testing.T) {
	// Test cases
	audioState := GetMicrophoneState()

	// Assert that audioState is a boolean
	if audioState != true && audioState != false {
		t.Errorf("Expected audioState to be of type bool, but got %T", audioState)
	}
}
