package media

import (
	"fmt"

	"github.com/cakeholeDC/on-air/appconfig"
	"github.com/cakeholeDC/on-air/logger"
	"github.com/fatih/color"
)

var log = logger.New("media")

func ShouldBeOn() bool {
	// returns true if ANY media device is enabled.
	return GetCameraState() || GetMicrophoneState()
}

func PrintMediaStates() {
	green := color.New(color.FgHiGreen)
	red := color.New(color.FgHiRed)

	cfg, _ := appconfig.GetConfig()

	if cfg.OnairEnableCamera {
		cameraState := GetCameraState()
		if cameraState {
			fmt.Print("video: ")
			green.Print("ON\n")
		} else {
			fmt.Print("video: ")
			red.Print("OFF\n")
		}
	} else {
		fmt.Print("video: ")
		red.Print("SERVICE DISABLED\n")
	}

	if cfg.OnairEnableMicrophone {
		microphoneState := GetMicrophoneState()
		if microphoneState {
			fmt.Print("audio: ")
			green.Print("ON\n")
		} else {
			fmt.Print("audio: ")
			red.Print("OFF\n")
		}
	} else {
		fmt.Print("audio: ")
		red.Print("SERVICE DISABLED\n")
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
