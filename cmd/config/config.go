package config

import (
	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/spf13/cobra"
	"github.com/zs5460/art"
)

var listFlag string = "list"
var endpointFlag string = "hass-endpoint"
var tokenFlag string = "hass-token"
var entityFlag string = "hass-entity"
var videoFlag string = "enable-video"
var audioFlag string = "enable-audio"

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the application",
	// TODO: Add encryption
	Long: art.String("onair.cfg") + "\nConfigure the application parameters such as the API endpoint, API key, and other settings.", //\n\nValues can encrypted on disk with the env vars 'IS_CFG_ENCRYPT' and 'CFG_CIPHER_KEY'",
	Run: func(cmd *cobra.Command, args []string) {
		// if no flags, return the help menu.
		if cmd.Flags().NFlag() == 0 {
			cmd.Help()
			return
		}

		// parse the flags
		list, _ := cmd.Flags().GetBool(listFlag)
		endpoint, _ := cmd.Flags().GetString(endpointFlag)
		token, _ := cmd.Flags().GetString(tokenFlag)
		device, _ := cmd.Flags().GetString(entityFlag)
		video, _ := cmd.Flags().GetString(videoFlag)
		audio, _ := cmd.Flags().GetString(audioFlag)

		// get the config
		cfg, _ := appconfig.GetConfig()

		// handle the flags
		if list {
			cfg.Print()
			return
		}
		if endpoint != "" {
			cfg.SetConfigValue("home_assistant_url", endpoint)
		}
		if token != "" {
			cfg.SetConfigValue("home_assistant_token", token)
		}
		if device != "" {
			cfg.SetConfigValue("home_assistant_entity", device)
		}
		if video != "" {
			cfg.SetConfigValue("onair_enable_camera", video)
		}
		if audio != "" {
			cfg.SetConfigValue("onair_enable_microphone", audio)
		}
		// always print the config
		cfg.Print()
	},
}

func init() {
	ConfigCmd.PersistentFlags().Bool(listFlag, false, "List the current configuration")
	ConfigCmd.PersistentFlags().String(endpointFlag, "", "Set the HASS endpoint url")
	ConfigCmd.PersistentFlags().String(tokenFlag, "", "Set the HASS API token")
	ConfigCmd.PersistentFlags().String(entityFlag, "", "Set the HASS entity name")
	ConfigCmd.PersistentFlags().String(videoFlag, "", "Enable the video trigger")
	ConfigCmd.PersistentFlags().String(audioFlag, "", "Enable the audio trigger")
}
