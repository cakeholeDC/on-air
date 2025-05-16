package cmd

import (
	"github.com/cakeholeDC/on-air/cmd/config"
	"github.com/spf13/cobra"
	"github.com/zs5460/art"
)

var rootCmd = cobra.Command{
	Use:   "onair",
	Short: "onair is a command line tool to control Home Assistant entities",
	Long:  art.String("onair") + "\n\033[1monair\033[0m is a command line tool built in go to control Home Assistant entities. It is designed to be simple and easy to use, with a focus on controlling media devices.",
}

func init() {
	rootCmd.AddCommand(config.ConfigCmd)
}

func Execute() {
	rootCmd.Execute()
}
