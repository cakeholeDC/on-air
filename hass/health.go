package hass

import (
	"io"
	"net/http"

	"encoding/json"

	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/cakeholeDC/on-air/common"
	"github.com/cakeholeDC/on-air/logger"
)

var log = logger.New("hass")

type HealthResponse struct {
	Message string `json:"message"`
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewHTTPClient() HTTPClient {
	return &http.Client{}
}

func Health(client HTTPClient, config *appconfig.AppConfig) (HealthResponse, error) {
	log.Info("Checking Home Assistant health...")
	// cfg, _ := appconfig.GetConfig()
	// req, err := http.NewRequest("GET", cfg.HomeAssistantURL+"/api/", nil)
	req, err := http.NewRequest("GET", config.HomeAssistantURL+"/api/", nil)
	if err != nil {
		log.Error(err.Error())
	}

	header := http.Header{
		// "Authorization": []string{"Bearer " + cfg.HomeAssistantToken},
		"Authorization": []string{"Bearer " + config.HomeAssistantToken},
		"Content-Type":  []string{"application/json"},
	}
	req.Header = header

	resp, err := client.Do(req)
	if err != nil {
		r := HealthResponse{}
		r.Message = "Failed to connect to Home Assistant"
		log.Error(err.Error())
		common.PrintJSON(r)
		return r, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response body: " + err.Error())
	}
	defer resp.Body.Close()

	r := HealthResponse{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Error("Failed to decode JSON response: " + err.Error())
	}

	common.PrintJSON(r)
	return r, nil
}
