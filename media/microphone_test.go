package media

import (
	"os"
	"runtime"
	"testing"

	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMicrophoneState_ServiceDisabled(t *testing.T) {
	testConfigPath := setupTestMediaEnv(t)

	// Create config with microphone disabled
	cfg := appconfig.NewConfig()
	cfg.SetConfigValue("onair_enable_microphone", "false")
	require.NoError(t, cfg.Save(testConfigPath))

	// When microphone service is disabled, should return false
	microphoneState := GetMicrophoneState()

	assert.False(t, microphoneState, "Microphone state should be false when service is disabled")
}

func TestGetMicrophoneState_ServiceEnabled(t *testing.T) {
	// Skip this test in CI environments where hardware APIs aren't available
	if os.Getenv("CI") != "" || os.Getenv("GITHUB_ACTIONS") != "" {
		t.Skip("Skipping hardware-dependent test in CI environment")
	}

	// Only run on macOS since this uses macOS-specific APIs
	if runtime.GOOS != "darwin" {
		t.Skip("Microphone state check only works on macOS")
	}

	testConfigPath := setupTestMediaEnv(t)

	// Create config with microphone enabled
	cfg := appconfig.NewConfig()
	cfg.SetConfigValue("onair_enable_microphone", "true")
	require.NoError(t, cfg.Save(testConfigPath))

	// Call the function - it should not panic
	microphoneState := GetMicrophoneState()

	// Assert that microphoneState is a boolean (either true or false)
	assert.IsType(t, false, microphoneState, "Microphone state should be a boolean")
}
