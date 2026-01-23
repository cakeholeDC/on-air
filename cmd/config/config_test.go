package config

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestConfigEnv creates a temporary directory and sets the environment variable
// to point to a test config file. This ensures we don't interfere with real config.
func setupTestConfigEnv(t *testing.T) string {
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

// executeCommand is a helper to execute a cobra command and capture output
func executeCommand(cmd *cobra.Command, args []string) (string, error) {
	// Reset all local flags to their default values before each test
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		flag.Value.Set(flag.DefValue)
		flag.Changed = false
	})

	// Also reset persistent flags
	cmd.PersistentFlags().VisitAll(func(flag *pflag.Flag) {
		flag.Value.Set(flag.DefValue)
		flag.Changed = false
	})

	// Reset command args first
	cmd.SetArgs(args)

	// Parse the flags to properly set the Changed status
	cmd.ParseFlags(args)

	// Capture stdout/stderr since the command uses fmt.Print directly
	oldStdout := os.Stdout
	oldStderr := os.Stderr

	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w

	// Execute command (this will parse flags again, but that's okay)
	err := cmd.Execute()

	// Restore stdout/stderr
	w.Close()
	os.Stdout = oldStdout
	os.Stderr = oldStderr

	// Read captured output
	output, _ := io.ReadAll(r)

	// Reset the command state for next test
	cmd.SetArgs([]string{})

	return string(output), err
}

func Test_CreateConfig_Success(t *testing.T) {
	testConfigPath := setupTestConfigEnv(t)

	// Ensure config doesn't exist before test
	_, err := os.Stat(testConfigPath)
	require.True(t, os.IsNotExist(err), "Config file should not exist before test")

	// Execute the create command
	output, err := executeCommand(ConfigCmd, []string{"--create"})

	// Assertions
	assert.NoError(t, err)
	assert.Contains(t, output, "Created a new configuration file")
	assert.Contains(t, output, testConfigPath)

	// Verify config file was actually created
	_, err = os.Stat(testConfigPath)
	assert.NoError(t, err, "Config file should exist after creation")
}

func Test_CreateConfig_FileAlreadyExists(t *testing.T) {
	testConfigPath := setupTestConfigEnv(t)

	// Create an existing config file
	cfg := appconfig.NewConfig()
	err := cfg.Save(testConfigPath)
	require.NoError(t, err, "Failed to create test config file")

	// Try to create config again
	output, err := executeCommand(ConfigCmd, []string{"--create"})

	// Assertions - should abort without error but with warning message
	assert.NoError(t, err)
	assert.Contains(t, output, "A configuration file already exists")
	assert.Contains(t, output, "Aborting to avoid overwriting")
}

func Test_CreateConfig_WithOtherFlags(t *testing.T) {
	setupTestConfigEnv(t)

	// Try to use --create with another flag (should be rejected)
	output, err := executeCommand(ConfigCmd, []string{"--create", "--list"})

	// Assertions
	assert.NoError(t, err)
	assert.Contains(t, output, "The '--create' flag cannot be used with any other flags")
}

func Test_ConfigList_NoConfigFile(t *testing.T) {
	testConfigPath := setupTestConfigEnv(t)

	// Ensure config doesn't exist
	_, err := os.Stat(testConfigPath)
	require.True(t, os.IsNotExist(err), "Config file should not exist")

	// Try to list config when it doesn't exist
	output, err := executeCommand(ConfigCmd, []string{"--list"})

	// Assertions - should fail gracefully with helpful message
	assert.NoError(t, err) // cobra command shouldn't error, but should print error message
	assert.Contains(t, output, "a configuration file is required")
	assert.Contains(t, output, "onair config --help")
}

func Test_ConfigList_Success(t *testing.T) {
	testConfigPath := setupTestConfigEnv(t)

	// Create a config with some values
	cfg := appconfig.NewConfig()
	cfg.SetConfigValue("home_assistant_url", "http://homeassistant.local:8123")
	cfg.SetConfigValue("home_assistant_token", "test-token-123")
	cfg.SetConfigValue("home_assistant_entity", "light.office")

	// List the config
	output, err := executeCommand(ConfigCmd, []string{"--list"})

	// Assertions
	assert.NoError(t, err)
	assert.Contains(t, output, testConfigPath)
	assert.Contains(t, output, "http://homeassistant.local:8123")
	assert.Contains(t, output, "test-token-123")
	assert.Contains(t, output, "light.office")
}

func Test_SetFlags_UpdatesConfig(t *testing.T) {
	testConfigPath := setupTestConfigEnv(t)

	// Create initial config
	cfg := appconfig.NewConfig()
	err := cfg.Save(testConfigPath)
	require.NoError(t, err)

	// Update values using flags
	output, err := executeCommand(ConfigCmd, []string{
		"--hass-endpoint", "http://new-endpoint:8123",
		"--hass-token", "new-token",
		"--hass-entity", "switch.office_light",
		"--enable-video", "true",
		"--enable-audio", "false",
	})

	// Assertions
	assert.NoError(t, err)
	assert.Contains(t, output, "http://new-endpoint:8123")
	assert.Contains(t, output, "new-token")
	assert.Contains(t, output, "switch.office_light")

	// Verify values were persisted to disk
	updatedCfg, err := appconfig.GetConfig()
	require.NoError(t, err)
	assert.Equal(t, "http://new-endpoint:8123", updatedCfg.HomeAssistantURL)
	assert.Equal(t, "new-token", updatedCfg.HomeAssistantToken)
	assert.Equal(t, "switch.office_light", updatedCfg.HomeAssistantEntity)
	assert.Equal(t, true, updatedCfg.OnairEnableCamera)
	assert.Equal(t, false, updatedCfg.OnairEnableMicrophone)
}

func Test_NoFlags_WithoutConfig_ShowsError(t *testing.T) {
	testConfigPath := setupTestConfigEnv(t)

	// Ensure no config exists
	_, err := os.Stat(testConfigPath)
	require.True(t, os.IsNotExist(err), "Config file should not exist")

	// Execute command with no flags
	output, err := executeCommand(ConfigCmd, []string{})

	// Assertions - when no flags and no config, it shows an error
	assert.NoError(t, err)
	assert.Contains(t, output, "a configuration file is required")
	assert.Contains(t, output, "onair config --help")
}
