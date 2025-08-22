package media

import (
	"strconv"

	mediaDevices "github.com/cakeholeDC/go-media-devices-state"
)

func GetMicrophoneState() bool {
	log.Debug("Checking microphone state...")
	var isMicrophoneOn bool
	var err error

	// If using forked package
	isMicrophoneOn, err = mediaDevices.IsMicrophoneOn(false)
	// // If NOT using FORK
	// silenceNSLog(func() {
	// 	isMicrophoneOn, err = mediaDevices.IsMicrophoneOn(false)
	// })

	if err != nil {
		println("Error", "err", err)
		log.Error("Error checking microphone state: " + err.Error())
		return false
	}
	log.Debug("Microphone state is: " + strconv.FormatBool(isMicrophoneOn))
	return isMicrophoneOn
}
