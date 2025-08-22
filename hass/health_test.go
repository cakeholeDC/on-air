package hass

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockClient is a mock implementation of the HTTP client
type MockClient struct {
	mock.Mock
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestHealth(t *testing.T) {
	mockClient := new(MockClient)
	mockConfig := &appconfig.AppConfig{
		HomeAssistantURL:   "http://localhost:8123",
		HomeAssistantToken: "test_token",
	}

	// Mock response body
	mockResponse := `{"message": "API Running."}`
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(mockResponse)),
	}

	// Mock the Do method
	mockClient.On("Do", mock.Anything).Return(resp, nil)

	// Call the Health function with the mock client
	healthResp, err := Health(mockClient, mockConfig)

	// Assert the returned HealthResponse
	assert.NoError(t, err)
	assert.Equal(t, "API Running.", healthResp.Message)

	// Assert that the mock client's Do method was called
	mockClient.AssertExpectations(t)
}
