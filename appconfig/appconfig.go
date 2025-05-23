package appconfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/cakeholeDC/on-air/common"
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

func init() {
	// TODO: remove all this.
	homedir, _ := os.UserHomeDir()
	fmt.Println("HOMEDIR=" + homedir)
	pwd, _ := os.Getwd()
	fmt.Println("PWD=" + pwd)
}

func (c *AppConfig) Print() {
	fmt.Println(art.String("onair.cfg"))
	common.PrintYAML(c)
}

func GetConfig() (*AppConfig, error) {
	// create a new config struct
	c := &AppConfig{}
	// read the file from disk
	data, err := common.ReadFileBlob(CFG_FILE_PATH)

	// check if the file exists
	if err != nil {
		// if the file is not found, create a new config file
		if strings.Contains(err.Error(), "no such file") {
			fmt.Println("Config file does not exist.")
			fmt.Printf("Creating config file: \033[1m%s\033[0m\n", CFG_FILE_PATH)
			// write the empty config to the file
			c.Save(CFG_FILE_PATH)
			return c, nil
		} else {
			return nil, err
		}
	}

	// parse the config into a struct
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *AppConfig) SetConfigValue(key string, value string) {
	switch key {
	case "home_assistant_url":
		c.HomeAssistantURL = value
	case "home_assistant_token":
		c.HomeAssistantToken = value
	case "home_assistant_entity":
		c.HomeAssistantEntity = value
	default:
		return
	}

	c.Save(CFG_FILE_PATH)
}

func (c *AppConfig) Save(filePath string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return common.WriteFileBlob(filePath, data)
}
