package hass

import (
	"fmt"
	"os"

	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/cakeholeDC/on-air/constants"

	"github.com/fatih/color"
)

func getCachePath() (string, error) {
	// Cache is not configurable by environment variables.
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Error(fmt.Sprintf("could not determine user home directory: %s", err))

		return "", nil
	}

	cfg, err := appconfig.GetConfig()
	if err != nil {
		log.Error(fmt.Sprintf("failed to get config: %s", err))

		return "", err
	}

	cacheConfigPath := fmt.Sprintf(
		"%s/%s/%s",
		userHomeDir,
		constants.ONAIR_DEFAULT_HOME_APP_DATA_SUFFIX,
		cfg.HomeAssistantEntity+".cache",
	)
	// return the config file path
	return cacheConfigPath, nil
}

// read cache from disk. If no cache, assume off.
func readCache() (bool, error) {
	cachePath, err := getCachePath()
	if err != nil {
		return false, err
	}

	log.Info("reading cache from " + cachePath)
	data, err := os.ReadFile(cachePath)
	if err != nil {
		// assume off if no cache
		return false, nil
	}

	log.Debug("cache data: " + string(data))
	log.Info("cache read successfully")
	if string(data) == "true" {
		return true, nil
	} else {
		return false, nil
	}
}

// write cache to disk
func writeCache(state bool) error {
	cachePath, err := getCachePath()
	if err != nil {
		return err
	}

	log.Info("writing cache to " + cachePath)
	log.Debug("cache value: " + fmt.Sprintf("%t", state))
	data := "false"
	if state {
		data = "true"
	}
	err = os.WriteFile(cachePath, []byte(data), 0644)
	if err != nil {
		return err
	}
	log.Info("cache written successfully")

	return nil
}

func PrintCache() {
	green := color.New(color.FgHiGreen)
	red := color.New(color.FgHiRed)
	cfg, err := appconfig.GetConfig()
	if err != nil {
		log.Error(fmt.Sprintf("failed to get config: %s", err))

		return
	}

	cacheState, err := readCache()
	if err != nil {
		log.Error(fmt.Sprintf("failed to read cache: %s", err))
	}

	fmt.Print(cfg.HomeAssistantEntity + ".cache: ")
	if cacheState {
		green.Println("ON")
	} else {
		red.Println("OFF")
	}
}
