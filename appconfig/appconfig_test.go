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
