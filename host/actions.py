import json
import os
import psutil

from api.actions import get_smartthings_device_status
from logger import write_log_file

TRIGGER_APPS = json.loads(os.getenv('TRIGGER_APPS'))
DEVICE_NAME = os.getenv('DEVICE_NAME')
STATUS_LOG = os.getenv('STATUS_LOG')

def discover_process_names():
    for proc in psutil.process_iter():
        print(proc.name())
        

def is_trigger_app_running():
    '''
    parses running processes and checkes if any TRIGGER_APPS are running

    returns BOOLEAN
    '''
    processes = []
    for proc in psutil.process_iter():
        try:
            processName = proc.name()
            if processName in TRIGGER_APPS:
                processes.append(processName)
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
        light_status = open(STATUS_LOG, "r").readlines()[0] == "True\n"
    else:
        print("STATUS_LOG not present")
        light_status = get_smartthings_device_status(DEVICE_NAME)
        # write to STATUS_LOG
        write_log_file(light_status)

    return light_status
