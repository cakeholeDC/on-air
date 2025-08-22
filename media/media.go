package media

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintMediaStates() {
	// Print the current states of media devices
	cameraState := GetCameraState()
	microphoneState := GetMicrophoneState()

	green := color.New(color.FgHiGreen)
	red := color.New(color.FgHiRed)

	if cameraState {
		fmt.Print("video: ")
		green.Print("ON\n")
	} else {
		fmt.Print("video: ")
		red.Print("OFF\n")
	}

	if microphoneState {
		fmt.Print("audio: ")
		green.Print("ON\n")
	} else {
		fmt.Print("audio: ")
		red.Print("OFF\n")
	}
}
