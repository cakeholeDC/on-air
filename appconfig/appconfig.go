package appconfig

import (
	"fmt"
	"strings"

	"github.com/cakeholeDC/on-air/common"
	"github.com/cakeholeDC/on-air/logger"
	"github.com/zs5460/art"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	HomeAssistantURL      string `yaml:"home_assistant_url"`
	HomeAssistantToken    string `yaml:"home_assistant_token"`
	HomeAssistantEntity   string `yaml:"home_assistant_entity"`
	OnairEnableCamera     bool   `yaml:"onair_enable_camera"`
	OnairEnableMicrophone bool   `yaml:"onair_enable_microphone"`
}

var CFG_FILE_PATH string = "onair.cfg"

var log = logger.New("config")

func init() {
	// TODO: remove all this.
	log.Error("🚨 This isn't an error. This is the init function of the config module. Make sure this gets removed.")
}

func (c *AppConfig) Print() {
	fmt.Println(art.String("onair.cfg"))
	common.PrintYAML(c)
}

func GetConfig() (*AppConfig, error) {
	// create a new config struct
	c := &AppConfig{}
	// read the file from disk
	log.Debug(fmt.Sprintf("reading config: %s", CFG_FILE_PATH))
	data, err := common.ReadFileBlob(CFG_FILE_PATH)

	// check if the file exists
	if err != nil {
		// if the file is not found, create a new config file
		if strings.Contains(err.Error(), "no such file") {
			// log the action
			log.Info("Config file does not exist.")
			log.Info(fmt.Sprintf("Creating config file: %s", CFG_FILE_PATH))
			// print for the user
			fmt.Println("Config file does not exist.")
			fmt.Printf("Creating config file: %s\n", CFG_FILE_PATH)
			// write the empty config to the file
			c.Save(CFG_FILE_PATH)
			return c, nil
		} else {
			log.Error(fmt.Sprintf("Error reading config file: %s", err))
			return nil, err
		}
	}

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

	c.Save(CFG_FILE_PATH)
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
