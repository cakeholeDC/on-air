package hass

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/cakeholeDC/on-air/common"
)

type HassEntity struct {
	EntityID   string `json:"entity_id"`
	State      string `json:"state"`
	Attributes struct {
		Icon         string `json:"icon"`
		FriendlyName string `json:"friendly_name"`
	} `json:"attributes"`
	LastChanged  string `json:"last_changed"`
	LastReported string `json:"last_reported"`
	LastUpdated  string `json:"last_updated"`
	Context      struct {
		ID       string      `json:"id"`
		ParentID interface{} `json:"parent_id"`
		UserID   string      `json:"user_id"`
	} `json:"context"`
}

var cfg, _ = appconfig.GetConfig()

func (e HassEntity) Print() {
	common.PrintYAML(e)
}

// GetEntity retrieves the state of a specific entity from Home Assistant
func GetEntity() HassEntity {
	return GetEntityWithClient(&http.Client{}, cfg.HomeAssistantURL, cfg.HomeAssistantEntity, cfg.HomeAssistantToken)
}

// GetEntityWithClient retrieves the state of a specific entity from Home Assistant using a custom HTTP client
func GetEntityWithClient(client HTTPClient, baseURL, entityID, token string) HassEntity {
	req, err := http.NewRequest("GET", baseURL+"/api/states/"+entityID, nil)
	if err != nil {
		log.Error(err.Error())
		return HassEntity{}
	}

	header := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Content-Type":  []string{"application/json"},
	}
	req.Header = header

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())
		return HassEntity{}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response body: " + err.Error())
		return HassEntity{}
	}
	defer resp.Body.Close()

	entity := HassEntity{}
	err = json.Unmarshal(body, &entity)
	if err != nil {
		log.Error("Failed to decode JSON response: " + err.Error())
		return HassEntity{}
	}
	return entity
}

func ToggleEntity() []HassEntity {
	return ToggleEntityWithClient(&http.Client{}, cfg.HomeAssistantURL, cfg.HomeAssistantEntity, cfg.HomeAssistantToken)
}

func ToggleEntityWithClient(client HTTPClient, baseURL, entityID, token string) []HassEntity {
	req, err := http.NewRequest("POST", baseURL+"/api/services/homeassistant/toggle", nil)
	req.Body = io.NopCloser(strings.NewReader(`{"entity_id": "` + entityID + `"}`))

	if err != nil {
		log.Error(err.Error())
		return []HassEntity{}
	}

	header := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Content-Type":  []string{"application/json"},
	}
	req.Header = header

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())
		return []HassEntity{}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response body: " + err.Error())
		return []HassEntity{}
	}
	defer resp.Body.Close()

	var entities []HassEntity
	err = json.Unmarshal(body, &entities)
	if err != nil {
		log.Error("Failed to decode JSON response: " + err.Error())
		return []HassEntity{}
	}
	return entities
}
