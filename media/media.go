package media

import (
	"fmt"

	"github.com/fatih/color"
)

func ShouldBeOn() bool {
	// returns true if ANY media device is enabled.
	return GetCameraState() || GetMicrophoneState()
}

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

// This function will silence NSLog output
// func silenceNSLog(f func()) error {
// 	// Save the original stderr
// 	oldStderr := syscall.Stderr

// 	// Open /dev/null
// 	devNull, err := os.OpenFile("/dev/null", os.O_WRONLY, 0)
// 	if err != nil {
// 		return err
// 	}
// 	defer devNull.Close()

// 	// Redirect stderr to /dev/null
// 	err = syscall.Dup2(int(devNull.Fd()), int(os.Stderr.Fd()))
// 	if err != nil {
// 		return err
// 	}

// 	// Call the noisy function
// 	f()

// 	// Restore stderr
// 	err = syscall.Dup2(oldStderr, int(os.Stderr.Fd()))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
