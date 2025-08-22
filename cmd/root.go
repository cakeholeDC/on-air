package cmd

import (
	"fmt"

	"github.com/cakeholeDC/on-air/cmd/config"
	"github.com/cakeholeDC/on-air/cmd/hass"
	"github.com/cakeholeDC/on-air/logger"
	"github.com/cakeholeDC/on-air/media"
	"github.com/spf13/cobra"
	"github.com/zs5460/art"
)

var log = logger.New("rootcmd")

var rootCmd = cobra.Command{
	Use:   "onair",
	Short: "onair is a command line tool to control Home Assistant entities",
	Long:  art.String("onair") + "\n\033[1monair\033[0m is a command line tool built in go to control Home Assistant entities. It is designed to be simple and easy to use, with a focus on controlling media devices.",
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
	rootCmd.AddCommand(hass.HassCmd)
	rootCmd.AddCommand(checkCmd)
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
