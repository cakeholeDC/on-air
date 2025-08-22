package media

import (
	"os"
	"strconv"
	"syscall"

	mediaDevices "github.com/cakeholeDC/go-media-devices-state"
	"github.com/cakeholeDC/on-air/logger"
)

var log = logger.New("media")

func GetCameraState() bool {
	log.Debug("Checking camera state...")
	var isCameraOn bool
	var err error

	// if using forked package
	isCameraOn, err = mediaDevices.IsCameraOn(false)
	// // if NOT using FORK
	// silenceNSLog(func() {
	// 	isCameraOn, err = mediaDevices.IsCameraOn(false)
	// })

	if err != nil {
		log.Error("Error checking camera state: " + err.Error())
		return false
	}
	log.Debug("Camera state is: " + strconv.FormatBool(isCameraOn))
	return isCameraOn
}

func silenceNSLog(f func()) error {
	// Save the original stderr
	oldStderr := syscall.Stderr

	// Open /dev/null
	devNull, err := os.OpenFile("/dev/null", os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer devNull.Close()

	// Redirect stderr to /dev/null
	err = syscall.Dup2(int(devNull.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		return err
	}

	// Call the noisy function
	f()

	// Restore stderr
	err = syscall.Dup2(oldStderr, int(os.Stderr.Fd()))
	if err != nil {
		return err
	}

	return nil
}
