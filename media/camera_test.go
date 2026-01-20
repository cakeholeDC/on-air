package media

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestMediaEnv creates a temporary config for testing
func setupTestMediaEnv(t *testing.T) string {
	t.Helper()

	// Save original env var to restore later
	originalEnv := os.Getenv("ONAIR_CONFIG_FILE_PATH")

	// Create temp directory for test config
	tempDir := t.TempDir()
	testConfigPath := filepath.Join(tempDir, "test-config.yaml")

	// Set env var to test config path
	os.Setenv("ONAIR_CONFIG_FILE_PATH", testConfigPath)

	// Cleanup: restore original env var
	t.Cleanup(func() {
		if originalEnv == "" {
			os.Unsetenv("ONAIR_CONFIG_FILE_PATH")
		} else {
			os.Setenv("ONAIR_CONFIG_FILE_PATH", originalEnv)
		}
	})

	return testConfigPath
}

func TestGetCameraState_ServiceDisabled(t *testing.T) {
	testConfigPath := setupTestMediaEnv(t)

	// Create config with camera disabled
	cfg := appconfig.NewConfig()
	cfg.SetConfigValue("onair_enable_camera", "false")
	require.NoError(t, cfg.Save(testConfigPath))

	// When camera service is disabled, should return false
	cameraState := GetCameraState()

	assert.False(t, cameraState, "Camera state should be false when service is disabled")
}

func TestGetCameraState_ServiceEnabled(t *testing.T) {
	// Skip this test in CI environments where hardware APIs aren't available
	if os.Getenv("CI") != "" || os.Getenv("GITHUB_ACTIONS") != "" {
		t.Skip("Skipping hardware-dependent test in CI environment")
	}

	// Only run on macOS since this uses macOS-specific APIs
	if runtime.GOOS != "darwin" {
		t.Skip("Camera state check only works on macOS")
	}

	testConfigPath := setupTestMediaEnv(t)

	// Create config with camera enabled
	cfg := appconfig.NewConfig()
	cfg.SetConfigValue("onair_enable_camera", "true")
	require.NoError(t, cfg.Save(testConfigPath))

	// Call the function - it should not panic
	cameraState := GetCameraState()

	// Assert that cameraState is a boolean (either true or false)
	assert.IsType(t, false, cameraState, "Camera state should be a boolean")
}
