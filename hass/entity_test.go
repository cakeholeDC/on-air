package hass

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEntity(t *testing.T) {
	// Mock Home Assistant API response
	mockResponse := `{
		"entity_id": "switch.mock",
		"state": "on",
		"attributes": {
			"icon": "mdi:lightbulb",
			"friendly_name": "mock switch"
		},
		"last_changed": "2023-01-01T12:00:00Z",
		"last_updated": "2023-01-01T12:00:00Z",
		"context": {
			"id": "abc123",
			"parent_id": null,
			"user_id": "user123"
		}
	}`

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/states/switch.mock", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Create a mock HTTP client that uses our test server
	client := &http.Client{}

	// Call the function with our mock server URL and test credentials
	entity, _ := GetEntityWithClient(client, server.URL, "switch.mock", "test-token")

	// Assert the response
	assert.Equal(t, "switch.mock", entity.EntityID)
	assert.Equal(t, "on", entity.State)
	assert.Equal(t, "mdi:lightbulb", entity.Attributes.Icon)
	assert.Equal(t, "mock switch", entity.Attributes.FriendlyName)
	assert.Equal(t, "abc123", entity.Context.ID)
	assert.Equal(t, "user123", entity.Context.UserID)
}

func TestToggleEntity(t *testing.T) {
	// Mock Home Assistant API response for toggle service
	mockResponse := `[]` // Toggle service typically returns an empty array

	// Track the current state - simulate a toggle behavior
	currentState := "off"

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/services/homeassistant/toggle" {
			// Handle toggle request
			assert.Equal(t, "POST", r.Method)
			assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

			// Read and verify the request body
			body, _ := io.ReadAll(r.Body)
			assert.Contains(t, string(body), `"entity_id": "switch.mock"`)

			// Toggle the state
			if currentState == "off" {
				currentState = "on"
			} else {
				currentState = "off"
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockResponse))
		} else if r.URL.Path == "/api/states/switch.mock" {
			// Handle state request - return current state
			stateResponse := `{
				"entity_id": "switch.mock",
				"state": "` + currentState + `",
				"attributes": {
					"icon": "mdi:lightbulb",
					"friendly_name": "mock switch"
				},
				"last_changed": "2023-01-01T12:00:00Z",
				"last_updated": "2023-01-01T12:00:00Z",
				"context": {
					"id": "abc123",
					"parent_id": null,
					"user_id": "user123"
				}
			}`
			assert.Equal(t, "GET", r.Method)
			assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(stateResponse))
		}
	}))
	defer server.Close()

	// Create a mock HTTP client
	client := &http.Client{}

	// Get initial state
	initialEntity, _ := GetEntityWithClient(client, server.URL, "switch.mock", "test-token")
	initialState := initialEntity.State

	// Call the toggle function
	ToggleEntityWithClient(client, server.URL, "switch.mock", "test-token")

	// Get the state after toggle
	newEntity, _ := GetEntityWithClient(client, server.URL, "switch.mock", "test-token")
	newState := newEntity.State

	// Assert the state has changed
	assert.NotEqual(t, initialState, newState)
	assert.Equal(t, "on", newState) // Should be "on" since we started with "off"
}

//nolint:funlen
func TestSetEntityState(t *testing.T) {
	// Track the current state - simulate state changes
	currentState := "off"

	// Create a test server that handles both service calls and state queries
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/services/switch/turn_on" {
			// Handle turn_on service call
			assert.Equal(t, "POST", r.Method)
			assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

			// Read and verify the request body
			body, _ := io.ReadAll(r.Body)
			assert.Contains(t, string(body), `"entity_id": "switch.mock"`)

			// Change state to on
			currentState = "on"

			// Service calls typically return empty array
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[]`))
		} else if r.URL.Path == "/api/services/switch/turn_off" {
			// Handle turn_off service call
			assert.Equal(t, "POST", r.Method)
			assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

			// Read and verify the request body
			body, _ := io.ReadAll(r.Body)
			assert.Contains(t, string(body), `"entity_id": "switch.mock"`)

			// Change state to off
			currentState = "off"

			// Service calls typically return empty array
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[]`))
		} else if r.URL.Path == "/api/states/switch.mock" {
			// Handle state query
			assert.Equal(t, "GET", r.Method)
			assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

			stateResponse := `{
				"entity_id": "switch.mock",
				"state": "` + currentState + `",
				"attributes": {
					"icon": "mdi:lightbulb",
					"friendly_name": "mock switch"
				},
				"last_changed": "2023-01-01T12:00:00Z",
				"last_updated": "2023-01-01T12:00:00Z",
				"context": {
					"id": "abc123",
					"parent_id": null,
					"user_id": "user123"
				}
			}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(stateResponse))
		}
	}))
	defer server.Close()

	// Create a mock HTTP client
	client := &http.Client{}

	// Test turning entity ON
	t.Run("TurnOn", func(t *testing.T) {
		// Reset to known state
		currentState = "off"

		// Get initial state
		initialEntity, _ := GetEntityWithClient(client, server.URL, "switch.mock", "test-token")
		assert.Equal(t, "off", initialEntity.State)

		// Call SetEntityState to turn on
		SetEntityStateWithClient(client, server.URL, "switch.mock", "test-token", true)

		// Verify state changed to on
		updatedEntity, _ := GetEntityWithClient(client, server.URL, "switch.mock", "test-token")
		assert.Equal(t, "on", updatedEntity.State)
		assert.Equal(t, "switch.mock", updatedEntity.EntityID)
	})

	// Test turning entity OFF
	t.Run("TurnOff", func(t *testing.T) {
		// Reset to known state
		currentState = "on"

		// Get initial state
		initialEntity, _ := GetEntityWithClient(client, server.URL, "switch.mock", "test-token")
		assert.Equal(t, "on", initialEntity.State)

		// Call SetEntityState to turn off
		SetEntityStateWithClient(client, server.URL, "switch.mock", "test-token", false)

		// Verify state changed to off
		updatedEntity, _ := GetEntityWithClient(client, server.URL, "switch.mock", "test-token")
		assert.Equal(t, "off", updatedEntity.State)
		assert.Equal(t, "switch.mock", updatedEntity.EntityID)
	})
}
