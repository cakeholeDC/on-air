package hass

import (
	"github.com/cakeholeDC/on-air/hass"
	"github.com/spf13/cobra"
	"github.com/zs5460/art"
)

var healthFlag string = "health"
var getHassEntityFlag string = "entity"
var toggleEntityFlag string = "toggle"

var HassCmd = &cobra.Command{
	Use:   "hass",
	Short: "Home Assistant integration",
	Long:  art.String("onair.hass") + "\nInteract with Home Assitant",
	Run: func(cmd *cobra.Command, args []string) {
		// if no flags, return the help menu.
		if cmd.Flags().NFlag() == 0 {
			cmd.Help()
			return
		}

		health, _ := cmd.Flags().GetBool(healthFlag)
		getHassEntity, _ := cmd.Flags().GetBool(getHassEntityFlag)
		toggleEntity, _ := cmd.Flags().GetBool(toggleEntityFlag)

		if health {
			client := hass.NewHTTPClient() // or construct as appropriate for your project
			hass.Health(client)
			return
		}

		if getHassEntity {
			// client := hass.NewHTTPClient() // or construct as appropriate for your project
			// fmt.Println("Listing devices...", client)
			hass.GetEntity().Print()
			return
		}

		if toggleEntity {
			hass.ToggleEntity()
			return
		}
	},
}

func init() {
	HassCmd.PersistentFlags().Bool(healthFlag, false, "Check Home Assistant health status")
	HassCmd.PersistentFlags().Bool(getHassEntityFlag, false, "Get Home Assistant entity details")
	HassCmd.PersistentFlags().Bool(toggleEntityFlag, false, "Toggle Home Assistant entity")
}
