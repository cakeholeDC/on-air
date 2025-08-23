package appconfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/cakeholeDC/on-air/common"
	"github.com/cakeholeDC/on-air/constants"
	"github.com/cakeholeDC/on-air/logger"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	HomeAssistantURL      string `yaml:"home_assistant_url"`
	HomeAssistantToken    string `yaml:"home_assistant_token"`
	HomeAssistantEntity   string `yaml:"home_assistant_entity"`
	OnairEnableCamera     bool   `yaml:"onair_enable_camera"`
	OnairEnableMicrophone bool   `yaml:"onair_enable_microphone"`
}

var log = logger.New("config")

func ReadConfigPath() string {
	// Read the environment variable for the config file path
	envConfigPath := os.Getenv("ONAIR_CONFIG_FILE_PATH")

	// If the environment variable is not set, use the default path
	if envConfigPath == "" {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Error(fmt.Sprintf("could not determine user home directory: %s", err))

			return ""
		}
		envConfigPath = fmt.Sprintf(
			"%s/%s/%s",
			userHomeDir,
			constants.ONAIR_DEFAULT_HOME_APP_DATA_SUFFIX,
			constants.ONAIR_DEFAULT_CONFIG_FILENAME,
		)
	}
	// return the config file path
	return envConfigPath
}


func NewConfig() *AppConfig {
	return &AppConfig{}
}

func (c *AppConfig) Print() {
	modName := "onair.cfg"
	termWidth, _, _ := common.GetTerminalSize()
	fmt.Println(strings.Repeat("*", termWidth))
	common.PrintArtStringIfFits(modName)
	fmt.Println(strings.Repeat("*", termWidth))
	fmt.Printf("Configuration File: %s\n", ReadConfigPath())
	fmt.Println(strings.Repeat("*", termWidth))
	common.PrintYAML(c)
}

func GetConfig() (*AppConfig, error) {
	// read the env variable.
	cfgPath := ReadConfigPath()

	// using the env var, read the file from disk
	log.Debug(fmt.Sprintf("reading config: %s", cfgPath))
	data, err := common.ReadFileBlob(cfgPath)

	// check if the file exists
	if err != nil {
		log.Error(fmt.Sprintf("Error reading config file: %s", err))

		return nil, err
	}

	// create a new config struct
	c := &AppConfig{}

	// parse the config into a struct
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Error(fmt.Sprintf("Error parsing config file: %s", err))

		return nil, err
	}

	return c, nil
}

func (c *AppConfig) SetConfigValue(key string, value string) {
	log.Info(fmt.Sprintf("setting config: %s=%s", key, value))
	switch key {
	case "home_assistant_url":
		c.HomeAssistantURL = value
	case "home_assistant_token":
		c.HomeAssistantToken = value
	case "home_assistant_entity":
		c.HomeAssistantEntity = value
	case "onair_enable_camera":
		if value == "true" {
			c.OnairEnableCamera = true
		} else {
			c.OnairEnableCamera = false
		}
	case "onair_enable_microphone":
		if value == "true" {
			c.OnairEnableMicrophone = true
		} else {
			c.OnairEnableMicrophone = false
		}
	default:
		return
	}

	c.Save(ReadConfigPath())
}

func (c *AppConfig) Save(filePath string) error {
	log.Debug(fmt.Sprintf("saving config: %s", filePath))
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Error(fmt.Sprintf("Error marshalling config: %s", err))

		return err
	}

	return common.WriteFileBlob(filePath, data)
}
