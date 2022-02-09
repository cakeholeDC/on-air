from asyncore import write
import json
import os
import psutil
from LightControl import loop, getDeviceStatus

VIDEO_APPS = json.loads(os.getenv('VIDEO_APPS'))
DEVICE_NAME = os.getenv('DEVICE_NAME')
STATUS_LOG = os.getenv('STATUS_LOG')

def isVideoAppRunning():
    processes = []
    for proc in psutil.process_iter():
        try:
            # Get process name & pid from process object.
            processName = proc.name()
            # processID = proc.pid
            # print(processName , ' ::: ', processID)
            if processName in VIDEO_APPS:
            # to find an app, guess it's name
            # if "Photo Booth" in processName:
                processes.append(processName)
        except (psutil.NoSuchProcess, psutil.AccessDenied, psutil.ZombieProcess):
            pass

    print(processes)
    return len(processes) > 0


def get_smartthings_status(device_name):
    return loop.run_until_complete( getDeviceStatus(device_name) )


def checkAndLogLightStatus():
    log_exists = os.path.exists(STATUS_LOG)

    if log_exists:
        # read from local log file
        light_status = open(STATUS_LOG, "r").readlines()[0] == "True"
    else:
        light_status = get_smartthings_status(DEVICE_NAME)
        # write to local log file
        write_log_file(light_status)

    return light_status

def write_log_file(status):
    f = open(STATUS_LOG, "w")
    f.write(str(status))
    f.write("\n")
    f.close()

if __name__ == '__main__':
    checkAndLogLightStatus()