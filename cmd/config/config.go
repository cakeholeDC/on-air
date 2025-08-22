package config

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/cakeholeDC/on-air/logger"
	"github.com/spf13/cobra"
	"github.com/zs5460/art"
)

var log = logger.New("configcmd")

// TODO: think about the flag terminology.
var listFlag string = "list"
var endpointFlag string = "hass-endpoint"
var tokenFlag string = "hass-token"
var entityFlag string = "hass-entity"
var videoFlag string = "enable-video"
var audioFlag string = "enable-audio"

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the application",
	Long:  art.String("onair.cfg") + "\nConfigure the application parameters such as the API endpoint, API key, and other settings.",
	Run: func(cmd *cobra.Command, args []string) {
		// if no flags, return the help menu.
		if cmd.Flags().NFlag() == 0 {
			cmd.Help()
			printEnvVars()
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
			log.Debug("--list: Listing the current configuration")
			cfg.Print()
			return
		}
		if endpoint != "" {
			log.Debug("--hass-endpoint: Setting Home Assistant endpoint to " + endpoint)
			cfg.SetConfigValue("home_assistant_url", endpoint)
		}
		if token != "" {
			log.Debug("--hass-token: Setting Home Assistant token to " + token)
			cfg.SetConfigValue("home_assistant_token", token)
		}
		if device != "" {
			log.Debug("--hass-entity: Setting Home Assistant entity to " + device)
			cfg.SetConfigValue("home_assistant_entity", device)
		}
		if video != "" {
			log.Debug("--enable-video: Setting video trigger to " + video)
			cfg.SetConfigValue("onair_enable_camera", video)
		}
		if audio != "" {
			log.Debug("--enable-audio: Setting audio trigger to " + audio)
			cfg.SetConfigValue("onair_enable_microphone", audio)
		}
		// always print the config
		cfg.Print()
	},
}

func printEnvVars() {
	// Print the environment variables
	fmt.Println("------------------------")
	fmt.Println("Environment Variables:")
	fmt.Println("------------------------")
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		userHomeDir = "$HOME"
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ENV_VAR\tDEFAULT\tDESCRIPTION\t")
	fmt.Fprintln(w, fmt.Sprintf("ONAIR_CONFIG_FILE_PATH\t%s/.config/onair/onair.cfg\t(optional) path to the config file", userHomeDir))
	fmt.Fprintln(w, "ONAIR_CONFIG_ENCRYPTION_KEY\tnull\t(optional) key to encrypt/decrypt the config file")
	w.Flush()
}

func init() {
	ConfigCmd.PersistentFlags().Bool(listFlag, false, "List the current configuration")
	ConfigCmd.PersistentFlags().String(endpointFlag, "", "Set the HASS endpoint url")
	ConfigCmd.PersistentFlags().String(tokenFlag, "", "Set the HASS API token")
	ConfigCmd.PersistentFlags().String(entityFlag, "", "Set the HASS entity name")
	ConfigCmd.PersistentFlags().String(videoFlag, "", "Enable the video trigger")
	ConfigCmd.PersistentFlags().String(audioFlag, "", "Enable the audio trigger")
}
