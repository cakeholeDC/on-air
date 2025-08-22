package hassCmd

import (
	"fmt"

	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/cakeholeDC/on-air/hass"
	"github.com/cakeholeDC/on-air/logger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/zs5460/art"
)

var log = logger.New("hasscmd")
var red = color.New(color.FgRed).Add(color.Bold)

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
			log.Debug("--health: Checking Home Assistant health status...")
			client := hass.NewHTTPClient() // or construct as appropriate for your project
			cfg, err := appconfig.GetConfig()
			if err != nil {
				log.Error(err.Error())
				return
			}
			hass.Health(client, cfg)
			return
		}

		if getHassEntity {
			log.Debug("--entity: Getting Home Assistant entity details...")
			entity, err := hass.GetEntity()
			if err != nil {
				fmt.Printf("Failed to get entity state: %s\n", red.Sprint(err))
			} else {
				PrintOnAirASCII(entity.State)
				entity.Print()
			}
			return
		}

		if toggleEntity {
			log.Debug("--toggle: Toggling Home Assistant entity...")
			entity := hass.ToggleEntity()
			PrintOnAirASCII(entity[0].State)
			return
		}
	},
}

func PrintOnAirASCII(entityState string) {
	if entityState == "on" {
		green := color.New(color.FgHiGreen)
		red := color.New(color.FgHiRed)
		green.Println("------------------")
		green.Print("| --- ")
		red.Print("ON AIR")
		green.Print(" --- |\n")
		green.Println("------------------")
	} else {
		gray := color.New(color.FgHiBlack)
		black := color.New(color.FgBlack)
		gray.Println("------------------")
		gray.Print("| --- ")
		black.Print("ON AIR")
		gray.Print(" --- |\n")
		gray.Println("------------------")
	}
}

func init() {
	HassCmd.PersistentFlags().Bool(healthFlag, false, "Check Home Assistant health status")
	HassCmd.PersistentFlags().Bool(getHassEntityFlag, false, "Get Home Assistant entity details")
	HassCmd.PersistentFlags().Bool(toggleEntityFlag, false, "Toggle Home Assistant entity")
}
