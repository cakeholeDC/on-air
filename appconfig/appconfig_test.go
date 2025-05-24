package appconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestEnv(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		os.Remove(CFG_FILE_PATH)
	})
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
	c := AppConfig{}
	c.SetConfigValue("home_assistant_url", "http://example.com")
	c.SetConfigValue("home_assistant_token", "token123")
	c.SetConfigValue("onair_enable_camera", "true")

	cfg, _ := GetConfig()
	assert.Equal(t, "token123", cfg.HomeAssistantToken, "Expected home_assistant_token to be 'token123'")
	assert.Equal(t, "http://example.com", cfg.HomeAssistantURL, "Expected home_assistant_url to be 'http://example.com'")
	assert.Equal(t, true, cfg.OnairEnableCamera, "Expected onair_enable_camera to be 'true'")
}

