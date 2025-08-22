package config

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/cakeholeDC/on-air/logger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/zs5460/art"
)

var log = logger.New("configcmd")
var code = color.New(color.FgGreen, color.BgHiBlack).Add(color.Bold)
var green = color.New(color.FgGreen).Add(color.Bold)

var about = `
A configuration file is required to store application secrets and settings.

# Triggers
- audio [if camera == on then turnDevice(on)]
- video [if microphone == on then turnDevice(on)]

# Integrations
- hass-endpoint [URL of the Home Assistant instance]
- hass-token [API token for Home Assistant]
- hass-entity [Entity ID to control]
`

// TODO: think about the flag terminology.
var createFlag string = "create"
var listFlag string = "list"
var endpointFlag string = "hass-endpoint"
var tokenFlag string = "hass-token"
var entityFlag string = "hass-entity"
var videoFlag string = "enable-video"
var audioFlag string = "enable-audio"

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the application",
	Long:  art.String("onair.cfg") + about + fmt.Sprintf("\nRun %s to create a configuration file.", code.Sprintf("onair config --%s", createFlag)),
	Run: func(cmd *cobra.Command, args []string) {
		// if no flags, return the help menu.
		if cmd.Flags().NFlag() == 0 {
			cmd.Help()
			// printEnvVars()
			return
		}

		// parse the flags
		create, _ := cmd.Flags().GetBool(createFlag)
		list, _ := cmd.Flags().GetBool(listFlag)
		endpoint, _ := cmd.Flags().GetString(endpointFlag)
		token, _ := cmd.Flags().GetString(tokenFlag)
		device, _ := cmd.Flags().GetString(entityFlag)
		video, _ := cmd.Flags().GetString(videoFlag)
		audio, _ := cmd.Flags().GetString(audioFlag)

		// if create, don't check the config yet.
		if create {
			// warn if any other flags are used
			if list || endpoint != "" || token != "" || device != "" || video != "" || audio != "" {
				fmt.Println("The '--create' flag cannot be used with any other flags.")
				return
			}
			// do not overwrite if the file exists.
			cfgPath := appconfig.ReadConfigPath()
			if _, err := os.Stat(cfgPath); err == nil {
				fmt.Printf("A configuration file already exists at: %s\n", green.Sprint(cfgPath))
				fmt.Println("Aborting to avoid overwriting the existing file.")
				return
			} else if !os.IsNotExist(err) {
				fmt.Printf("Error checking configuration file: %v\n", err)
				return
			}

			// create a new config
			cfg := appconfig.NewConfig()
			cfg.Save(appconfig.ReadConfigPath())
			fmt.Printf("Created a new configuration file at: %s\n", green.Sprint(appconfig.ReadConfigPath()))
			return
		}

		// otherwise, check for a config
		cfg, err := appconfig.GetConfig()
		if err != nil {
			// fail if there is no config
			log.Error(fmt.Sprintf("Error getting config: %s", err))
			fmt.Printf("a configuration file is required for this application. please run %s for more info\n", code.Sprint("onair config --help"))
			return
		}

		// then handle the flags
		if list {
			log.Debug(fmt.Sprintf("--%s: Listing the current configuration", listFlag))
			cfg.Print()
			return
		}
		if endpoint != "" {
			log.Debug(fmt.Sprintf("--%s: Setting Home Assistant endpoint to %s", endpointFlag, endpoint))
			cfg.SetConfigValue("home_assistant_url", endpoint)
		}
		if token != "" {
			log.Debug(fmt.Sprintf("--%s: Setting Home Assistant token to %s", tokenFlag, token))
			cfg.SetConfigValue("home_assistant_token", token)
		}
		if device != "" {
			log.Debug(fmt.Sprintf("--%s: Setting Home Assistant entity to %s", entityFlag, device))
			cfg.SetConfigValue("home_assistant_entity", device)
		}
		if video != "" {
			log.Debug(fmt.Sprintf("--%s: Setting video trigger to %s", videoFlag, video))
			cfg.SetConfigValue("onair_enable_camera", video)
		}
		if audio != "" {
			log.Debug(fmt.Sprintf("--%s: Setting audio trigger to %s", audioFlag, audio))
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
	ConfigCmd.PersistentFlags().Bool(createFlag, false, "Create a new configuration file")
	ConfigCmd.PersistentFlags().Bool(listFlag, false, "List the current configuration")
	ConfigCmd.PersistentFlags().String(endpointFlag, "", "Set the HASS endpoint url")
	ConfigCmd.PersistentFlags().String(tokenFlag, "", "Set the HASS API token")
	ConfigCmd.PersistentFlags().String(entityFlag, "", "Set the HASS entity name")
	ConfigCmd.PersistentFlags().String(videoFlag, "", "Enable the video trigger")
	ConfigCmd.PersistentFlags().String(audioFlag, "", "Enable the audio trigger")
}
