package hass

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cakeholeDC/on-air/constants"
)

func setupTestCache(t *testing.T) (string, func()) {
	t.Helper()

	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "onair-cache-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Store original home dir and config path
	originalHome, _ := os.UserHomeDir()
	originalConfigPath := os.Getenv("ONAIR_CONFIG_FILE_PATH")

	// Set HOME to temp directory
	os.Setenv("HOME", tmpDir)

	// Create the cache directory structure
	cacheDir := filepath.Join(tmpDir, constants.ONAIR_DEFAULT_HOME_APP_DATA_SUFFIX)
	err = os.MkdirAll(cacheDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create cache dir: %v", err)
	}

	// Create a temporary config file with test entity
	configFile, err := os.CreateTemp(cacheDir, "onair-test-config-*.cfg")
	if err != nil {
		t.Fatalf("Failed to create temp config file: %v", err)
	}

	// Write minimal config with entity name for testing
	configContent := []byte("home_assistant_entity: test_entity\n")
	if err := os.WriteFile(configFile.Name(), configContent, 0644); err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	// Set config path environment variable
	os.Setenv("ONAIR_CONFIG_FILE_PATH", configFile.Name())

	// Return cleanup function
	cleanup := func() {
		os.Setenv("HOME", originalHome)
		if originalConfigPath != "" {
			os.Setenv("ONAIR_CONFIG_FILE_PATH", originalConfigPath)
		} else {
			os.Unsetenv("ONAIR_CONFIG_FILE_PATH")
		}
		os.RemoveAll(tmpDir)
	}

	return tmpDir, cleanup
}

func TestGetCachePath(t *testing.T) {
	tmpDir, cleanup := setupTestCache(t)
	defer cleanup()

	expectedPath := filepath.Join(tmpDir, constants.ONAIR_DEFAULT_HOME_APP_DATA_SUFFIX, "test_entity.cache")
	actualPath, err := getCachePath()
	if err != nil {
		t.Fatalf("getCachePath() error = %v, want nil", err)
	}

	if actualPath != expectedPath {
		t.Errorf("getCachePath() = %v, want %v", actualPath, expectedPath)
	}
}

func TestReadCache_NoFile(t *testing.T) {
	_, cleanup := setupTestCache(t)
	defer cleanup()

	// When no cache file exists, readCache should return false with no error
	state, err := readCache()

	if err != nil {
		t.Errorf("readCache() error = %v, want nil", err)
	}

	if state != false {
		t.Errorf("readCache() state = %v, want false", state)
	}
}

func TestReadCache_True(t *testing.T) {
	_, cleanup := setupTestCache(t)
	defer cleanup()

	// Write "true" to cache file
	cachePath, err := getCachePath()
	if err != nil {
		t.Fatalf("Failed to get cache path: %v", err)
	}
	err = os.WriteFile(cachePath, []byte("true"), 0644)
	if err != nil {
		t.Fatalf("Failed to write test cache file: %v", err)
	}

	state, err := readCache()

	if err != nil {
		t.Errorf("readCache() error = %v, want nil", err)
	}

	if state != true {
		t.Errorf("readCache() state = %v, want true", state)
	}
}

func TestReadCache_False(t *testing.T) {
	_, cleanup := setupTestCache(t)
	defer cleanup()

	// Write "false" to cache file
	cachePath, err := getCachePath()
	if err != nil {
		t.Fatalf("Failed to get cache path: %v", err)
	}
	err = os.WriteFile(cachePath, []byte("false"), 0644)
	if err != nil {
		t.Fatalf("Failed to write test cache file: %v", err)
	}

	state, err := readCache()

	if err != nil {
		t.Errorf("readCache() error = %v, want nil", err)
	}

	if state != false {
		t.Errorf("readCache() state = %v, want false", state)
	}
}

func TestReadCache_InvalidValue(t *testing.T) {
	_, cleanup := setupTestCache(t)
	defer cleanup()

	// Write an invalid value to cache file
	cachePath, err := getCachePath()
	if err != nil {
		t.Fatalf("Failed to get cache path: %v", err)
	}
	err = os.WriteFile(cachePath, []byte("invalid"), 0644)
	if err != nil {
		t.Fatalf("Failed to write test cache file: %v", err)
	}

	state, err := readCache()

	if err != nil {
		t.Errorf("readCache() error = %v, want nil", err)
	}

	// Any value other than "true" should be treated as false
	if state != false {
		t.Errorf("readCache() state = %v, want false for invalid value", state)
	}
}

func TestWriteCache_True(t *testing.T) {
	_, cleanup := setupTestCache(t)
	defer cleanup()

	err := writeCache(true)
	if err != nil {
		t.Errorf("writeCache(true) error = %v, want nil", err)
	}

	// Verify the file was written correctly
	cachePath, err := getCachePath()
	if err != nil {
		t.Fatalf("Failed to get cache path: %v", err)
	}
	data, err := os.ReadFile(cachePath)
	if err != nil {
		t.Fatalf("Failed to read cache file: %v", err)
	}

	if strings.TrimSpace(string(data)) != "true" {
		t.Errorf("Cache file content = %v, want 'true'", string(data))
	}
}

func TestWriteCache_False(t *testing.T) {
	_, cleanup := setupTestCache(t)
	defer cleanup()

	err := writeCache(false)
	if err != nil {
		t.Errorf("writeCache(false) error = %v, want nil", err)
	}

	// Verify the file was written correctly
	cachePath, err := getCachePath()
	if err != nil {
		t.Fatalf("Failed to get cache path: %v", err)
	}
	data, err := os.ReadFile(cachePath)
	if err != nil {
		t.Fatalf("Failed to read cache file: %v", err)
	}

	if strings.TrimSpace(string(data)) != "false" {
		t.Errorf("Cache file content = %v, want 'false'", string(data))
	}
}

func TestWriteAndReadCache_Roundtrip(t *testing.T) {
	_, cleanup := setupTestCache(t)
	defer cleanup()

	testCases := []struct {
		name  string
		state bool
	}{
		{"write and read true", true},
		{"write and read false", false},
		{"toggle from false to true", true},
		{"toggle from true to false", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Write cache
			err := writeCache(tc.state)
			if err != nil {
				t.Errorf("writeCache(%v) error = %v, want nil", tc.state, err)
			}

			// Read cache
			state, err := readCache()
			if err != nil {
				t.Errorf("readCache() error = %v, want nil", err)
			}

			if state != tc.state {
				t.Errorf("readCache() = %v, want %v", state, tc.state)
			}
		})
	}
}

func TestWriteCache_DirectoryDoesNotExist(t *testing.T) {
	// Create temp dir but don't create the cache directory structure
	tmpDir, err := os.MkdirTemp("", "onair-cache-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalHome, _ := os.UserHomeDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	// This should fail because the directory doesn't exist
	err = writeCache(true)
	if err == nil {
		t.Error("writeCache() error = nil, want error when directory doesn't exist")
	}
}
