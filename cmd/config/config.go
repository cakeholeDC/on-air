package config

import (
	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/spf13/cobra"
	"github.com/zs5460/art"
)

var listFlag string = "list"

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the application",
	Long:  art.String("onair.cfg") + "\nConfigure the application parameters such as the API endpoint, API key, and other settings. \n\nValues can encrypted on disk with the env vars 'IS_CFG_ENCRYPT' and 'CFG_CIPHER_KEY'",
	Run: func(cmd *cobra.Command, args []string) {
		// if no flags, return the help menu.
		if cmd.Flags().NFlag() == 0 {
			cmd.Help()
			return
		}

		// parse the flags
		list, _ := cmd.Flags().GetBool(listFlag)

		// get the config
		cfg, _ := appconfig.GetConfig()

		// handle the flags
		if list {
			cfg.Print()
			return
		}
	},
}

func init() {
	ConfigCmd.PersistentFlags().Bool(listFlag, false, "List the current configuration")
}
