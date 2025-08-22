package media

import (
	"strconv"

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
