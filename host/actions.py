import json
import os
import psutil

from api.actions import get_smartthings_device_status
from logger import write_light_status

TRIGGER_APPS = json.loads(os.getenv('TRIGGER_APPS'))
LIGHT_STATUS = os.getenv('LIGHT_STATUS')

def discover_process_names():
    for proc in psutil.process_iter():
        print(proc.name())


def is_trigger_app_running():
    '''
    parses running processes and checks if any TRIGGER_APPS are running

    returns BOOLEAN
    '''
    processes = []
    for proc in psutil.process_iter():
        try:
            process_name = proc.name()
            if process_name in TRIGGER_APPS:
                processes.append(process_name)
        except (psutil.NoSuchProcess, psutil.AccessDenied, psutil.ZombieProcess):
            pass

    # print("Open Trigger Apps:", processes)
    return len(processes) > 0


def get_and_log_device_status():
    '''
    returns BOOLEAN

    Parses STATUS_LOG for device status.

    If STATUS_LOG is missing, fetches api.device.status and writes result to STATUS_LOG.
    '''
    log_exists = os.path.exists(STATUS_LOG)

    if log_exists:
        # read from STATUS_LOG
        light_status = open(STATUS_LOG, "r").readlines()[0] == "True\n" # pylint: disable=[W1514, R1732]
    else:
        print("STATUS_LOG not present")
        light_status = get_smartthings_device_status()
        # write to STATUS_LOG
        write_log_file(light_status)

    return light_status
