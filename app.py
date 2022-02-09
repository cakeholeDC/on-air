import os
from LightControl import lightsOn, lightsOff, lights_object, loop
from hostActions import isVideoAppRunning, checkAndLogLightStatus

DEVICE_NAME = os.getenv('DEVICE_NAME')
STATUS_LOG = os.getenv('STATUS_LOG')


if __name__ == '__main__':
    webcamStatus = isVideoAppRunning()
    print(f"isAppOpen = {webcamStatus}")

    lightStatus = checkAndLogLightStatus()
    print(f"isLightOn = {lightStatus}")

    hamster = True
    if webcamStatus and lightStatus:
        pass
        # do nothing
    if not webcamStatus and not lightStatus:
        pass
    elif webcamStatus and not lightStatus:
        # turn light on
        loop.run_until_complete( lightsOn(lights_object, DEVICE_NAME) )
        # write to local log file
        f = open(STATUS_LOG, "w")
        f.write(str(True))
        f.close()

        print("Webcam: ON AIR")
    elif not webcamStatus and lightStatus:
        # turn light off
        loop.run_until_complete( lightsOff(lights_object, DEVICE_NAME) )
        # write to local log file
        f = open(STATUS_LOG, "w")
        f.write(str(False))
        f.close()

        print("Webcam: OFF AIR")
    else: # not webcamStatus and not lightStatus
        pass