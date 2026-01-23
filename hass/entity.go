package hass

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/cakeholeDC/on-air/common"
	"github.com/fatih/color"
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

func (e HassEntity) Print() {
	common.PrintYAML(e)
}

// GetEntity retrieves the state of a specific entity from Home Assistant
func GetEntity() (HassEntity, error) {
	cfg, err := appconfig.GetConfig()
	if err != nil {
		log.Error(err.Error())

		return HassEntity{}, err
	}

	return GetEntityWithClient(&http.Client{}, cfg.HomeAssistantURL, cfg.HomeAssistantEntity, cfg.HomeAssistantToken)
}

// GetEntityWithClient retrieves the state of a specific entity from Home Assistant using a custom HTTP client
func GetEntityWithClient(client HTTPClient, baseURL, entityID, token string) (HassEntity, error) {
	req, err := http.NewRequest(http.MethodGet, baseURL+"/api/states/"+entityID, nil)
	if err != nil {
		log.Error(err.Error())

		return HassEntity{}, err
	}

	header := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Content-Type":  []string{"application/json"},
	}
	req.Header = header

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())

		return HassEntity{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response body: " + err.Error())

		return HassEntity{}, err
	}
	defer resp.Body.Close()

	entity := HassEntity{}
	err = json.Unmarshal(body, &entity)
	if err != nil {
		log.Error("Failed to decode JSON response: " + err.Error())

		return HassEntity{}, err
	}

	return entity, nil
}

func ToggleEntity() []HassEntity {
	cfg, err := appconfig.GetConfig()
	if err != nil {
		log.Error(err.Error())

		return []HassEntity{}
	}

	return ToggleEntityWithClient(&http.Client{}, cfg.HomeAssistantURL, cfg.HomeAssistantEntity, cfg.HomeAssistantToken)
}

func ToggleEntityWithClient(client HTTPClient, baseURL, entityID, token string) []HassEntity {
	log.Debug("Toggling entity: " + entityID)
	req, err := http.NewRequest(http.MethodPost, baseURL+"/api/services/homeassistant/toggle", nil)
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

	log.Debug("Toggled entity: " + entityID)
	writeCache(entities[0].State == "on")

	return entities
}

func SetEntityState(state bool, skipCache bool) ([]HassEntity, error) {
	// read the cache to see the last known state
	cacheState, err := readCache()
	if err != nil {
		log.Error(fmt.Sprintf("failed to read cache: %s", err))
	}

	// if the desired state matches the cached state, do nothing (unless skipCache is true)
	if !skipCache && cacheState == state {
		log.Info("Entity state is already " +
			map[bool]string{true: "on", false: "off"}[state] +
			" according to cache; no action taken")

		return []HassEntity{}, nil
	}

	// otherwise, set the state and rewrite the cache
	cfg, err := appconfig.GetConfig()
	if err != nil {
		log.Error(err.Error())

		return []HassEntity{}, err
	}

	// sets the state and rewrites the cache
	return SetEntityStateWithClient(
		&http.Client{},
		cfg.HomeAssistantURL,
		cfg.HomeAssistantEntity,
		cfg.HomeAssistantToken,
		state,
	)
}

// SetEntityStateWithClient sets the state of a specific entity from Home Assistant using a custom HTTP client
func SetEntityStateWithClient(client HTTPClient, baseURL, entityID, token string, state bool) ([]HassEntity, error) {
	log.Info("Setting entity state: " + entityID + " to " + map[bool]string{true: "on", false: "off"}[state])

	url := ""
	if state {
		url = baseURL + "/api/services/switch/turn_on"
	} else {
		url = baseURL + "/api/services/switch/turn_off"
	}

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Error(err.Error())

		return []HassEntity{}, err
	}
	if !state {
		req.Body = io.NopCloser(strings.NewReader(`{"entity_id": "` + entityID + `"}`))
	} else {
		req.Body = io.NopCloser(strings.NewReader(`{"entity_id": "` + entityID + `"}`))
	}

	header := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Content-Type":  []string{"application/json"},
	}
	req.Header = header

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())

		return []HassEntity{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response body: " + err.Error())

		return []HassEntity{}, err
	}
	defer resp.Body.Close()

	var entities []HassEntity
	err = json.Unmarshal(body, &entities)
	if err != nil {
		log.Error("Failed to decode JSON response: " + err.Error())

		return []HassEntity{}, err
	}

	for _, ent := range entities {
		if ent.EntityID == entityID {
			log.Info("Set entity state: " + entityID + " to " + ent.State)

			break
		}
	}

	// write the new state to the cache
	writeCache(state)

	return entities, nil
}

func (e HassEntity) PrintState() {
	green := color.New(color.FgHiGreen)
	red := color.New(color.FgHiRed)

	fmt.Printf("%s: ", e.EntityID)
	if e.State == "on" {
		green.Print("ON\n")
	} else {
		red.Print("OFF\n")
	}

	log.Info(fmt.Sprintf("%s is: %s", e.EntityID, e.State))
}
