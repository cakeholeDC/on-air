package appconfig

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TEST_CONFIG_FILE string

func setupTestEnv(t *testing.T) {
	t.Helper()
	initEnv := os.Getenv("ONAIR_CONFIG_FILE_PATH") // This may not be necessary

	file, err := os.CreateTemp("", "onair-test-config-*.cfg")
	if err != nil {
		t.Fatalf("Failed to create temporary config file: %v", err)
	}
	TEST_CONFIG_FILE = file.Name()
	os.Setenv("ONAIR_CONFIG_FILE_PATH", TEST_CONFIG_FILE)
	t.Cleanup(func() {
		os.Unsetenv("ONAIR_CONFIG_FILE_PATH")
		os.Remove(TEST_CONFIG_FILE)
		os.Setenv("ONAIR_CONFIG_FILE_PATH", initEnv) // This may not be necessary.
	})
}

func TestReadConfigPath(t *testing.T) {
	setupTestEnv(t)
	// Tests that the config file path is read correctly from the environment variable
	cfgPath := readConfigPath()
	assert.Equal(t, TEST_CONFIG_FILE, cfgPath, fmt.Sprintf("Expected config file path to be '%s'", TEST_CONFIG_FILE))
}

func TestGetConfig(t *testing.T) {
	setupTestEnv(t)
	// Tests that the config file is read correctly
	cfg, err := GetConfig()
	if err != nil {
		t.Errorf("Error getting config: %v", err)
	}
	assert.Equal(t, "", cfg.HomeAssistantURL, "Expected home_assistant_url to be ''")
}

func TestSetConfigValue(t *testing.T) {
	setupTestEnv(t)
	// Tests that a config value is properly set, and written to disk
	c, _ := GetConfig()
	c.SetConfigValue("home_assistant_url", "test.com")
	c.SetConfigValue("home_assistant_token", "test_token")
	c.SetConfigValue("onair_enable_camera", "true")

	cfg, _ := GetConfig()
	assert.Equal(t, "test_token", cfg.HomeAssistantToken, "Expected home_assistant_token to be 'test_token'")
	assert.Equal(t, "test.com", cfg.HomeAssistantURL, "Expected home_assistant_url to be 'test.com'")
	assert.Equal(t, true, cfg.OnairEnableCamera, "Expected onair_enable_camera to be 'true'")
}

