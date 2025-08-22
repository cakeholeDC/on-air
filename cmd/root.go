package cmd

// TODO: ensure all this works
// - onair		 	  checks if the light should be on, and acts accordingly
// - onair --on 	  turns on the light
// - onair --off 	  turns off the light
// - onair --status   checks the status of the triggers and devices
// - onair config ... manages configuration
// - onair hass ...   interacts with home assistant directly

import (
	"fmt"

	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/cakeholeDC/on-air/cmd/config"
	"github.com/cakeholeDC/on-air/cmd/hassCmd"
	"github.com/cakeholeDC/on-air/hass"
	"github.com/cakeholeDC/on-air/logger"
	"github.com/cakeholeDC/on-air/media"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/zs5460/art"
)

var about string = `
onair is a command line tool to control home assistant entities.

this tool will check if the camera and/or microphone is active and toggle a device accordingly. 

run with flags to control the device directly.
`

var log = logger.New("rootcmd")
var red = color.New(color.FgRed).Add(color.Bold)

var onFlag string = "on"
var offFlag string = "off"
var checkFlag string = "status"

var rootCmd = cobra.Command{
	Use:   "onair",
	Short: "onair is a command line tool to control home assistant entities",
	// 	Long:  art.String("onair") + "\n\033[1monair\033[0m is a command line tool built in go to control Home Assistant entities. It is designed to be simple and easy to use, with a focus on controlling media devices.",
	Long: art.String("onair") + about,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 {
			// if the command is run without any flags and there is no configuration file present, show help
			cfg, err := appconfig.GetConfig()
			if err != nil || cfg == nil {
				cmd.Help()
				red.Println("\nonair requires a configuration file. Run 'onair config --help' for more information.")
				return
			}
			// otherwise, run in 'fast mode' - check if it should be on and act accordingly
			//? TODO: should fast mode be configurable?
			shouldBeOn := media.ShouldBeOn()
			if shouldBeOn {
				hass.SetEntityState(true)
				hassCmd.PrintOnAirASCII("on")
			} else {
				hass.SetEntityState(false)
				hassCmd.PrintOnAirASCII("off")
			}
			return
		}

		turnOn, _ := cmd.Flags().GetBool(onFlag)
		turnOff, _ := cmd.Flags().GetBool(offFlag)
		check, _ := cmd.Flags().GetBool(checkFlag)

		if turnOn {
			log.Info("device turned ON manually")
			_, err := hass.SetEntityState(true)
			if err == nil {
				hassCmd.PrintOnAirASCII("on")
			} else {
				fmt.Printf("Failed to turn on the device: %s\n", red.Sprint(err))
			}
		} else if turnOff {
			log.Info("device turned OFF manually")
			_, err := hass.SetEntityState(false)
			if err == nil {
				hassCmd.PrintOnAirASCII("off")
			} else {
				fmt.Printf("Failed to turn off the device: %s\n", red.Sprint(err))
			}
		} else if check {
			log.Info("running status checks...")
			media.PrintMediaStates()
			entity, err := hass.GetEntity()
			if err != nil {
				fmt.Printf("Failed to get entity state: %s\n", red.Sprint(err))
			} else {
				entity.PrintState()
			}
		} else {
			log.Warn("No valid flag provided")
		}
	},
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run onair checks",
	Long:  art.String("onair.check") + "\nRun onair checks",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("--check: Running onair checks")
		// fmt.Println("Running onair checks...")
		cameraState := media.GetCameraState()
		microphoneState := media.GetMicrophoneState()
		// fmt.Printf("Camera state: %s\n", strconv.FormatBool(cameraState))
		// fmt.Printf("Microphone state: %s\n", strconv.FormatBool(microphoneState))

		if cameraState || microphoneState {
			fmt.Println("At least one media device is active.")
			// TURN ON DEVICE
		} else {
			fmt.Println("No media devices are active.")
			// TURN OFF DEVICE
		}
	},
}

func init() {
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(hassCmd.HassCmd)
	rootCmd.AddCommand(checkCmd)

	rootCmd.PersistentFlags().BoolP(onFlag, "o", false, "Turn on the device")
	rootCmd.PersistentFlags().BoolP(offFlag, "f", false, "Turn off the device")
	rootCmd.PersistentFlags().BoolP(checkFlag, "c", false, "Check the status of the media triggers")
}

func Execute() {
	// TODO: implement "fast mode" where you invoke the binary and the light toggles
	fastMode := false
	if !fastMode {
		rootCmd.Execute()
	} else {
		log.Debug("Fast mode enabled, skipping help output")
		// do the thing
	}
}
