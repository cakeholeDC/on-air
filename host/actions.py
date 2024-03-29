import json
import os
import subprocess
import psutil

from api.actions import get_smartthings_device_status
from logger import write_camera_status, write_light_status

TRIGGER_APPS = json.loads(os.getenv('TRIGGER_APPS'))
LIGHT_STATUS = os.getenv('LIGHT_STATUS')
CAMERA_STATUS = os.getenv('CAMERA_STATUS')

def discover_process_names():
    '''
    iterates and prints running processes on the host
    '''
    for proc in psutil.process_iter():
        print(proc.name())

def run_vdc_assistant_shell_script():
    '''
    reads VDCAssistant_Power_State entries for the last 5 minutes from UDCExtension.PowerLog

    returns "On", "Off", or ""
    '''
    # TODO: can we pass n minutes to the script? # pylint: disable=[W0511]
    output = subprocess.check_output("./host/parse_vdc_assistant_power_log.sh", shell=True)
    return output.decode("utf-8")[:-2]

def transform_vdc_assistant_power_state():
    '''
    transforms VDCAssistant_Power_State to BOOLEAN, or None

    Allowed Responses: True, False, or None
    '''
    latest_log_status = run_vdc_assistant_shell_script()
    if len(latest_log_status) > 0:
        return latest_log_status == "On"

    return None


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

def get_and_log_vdc_status():
    '''
    Takes vdc_assistant_power_state boolean, and checks it against the local status logs to determine
    whether light should be on, or if the local log status needs to be updated.

    returns Boolean indicating if VDC Assistant is Active
    '''

    vdc_status = transform_vdc_assistant_power_state()

    # catch both True and False
    if vdc_status is not None: # this means the status changed.
        # write it
        camera_status = vdc_status
        write_camera_status(camera_status)
    else: # this means no log entry in the last 5 minutes, therefore the status didn't change.

        # so check if we've logged status before
        vdc_log_exists = os.path.exists(CAMERA_STATUS)

        # does the log file exist?
        if vdc_log_exists:
            # if so, read the existing status as camera status
            camera_status = open(CAMERA_STATUS, "r").readlines()[0] == "True\n" # pylint: disable=[W1514, R1732]
        else:
            # if not, write the log assuming the camera is off and wait for PowerOn event
            camera_status = False
            write_camera_status(camera_status)

    return camera_status

def get_and_log_device_status():
    '''
    returns BOOLEAN

    Parses LIGHT_STATUS for device status.

    If LIGHT_STATUS is missing, fetches api.device.status and writes result to LIGHT_STATUS.
    '''
    log_exists = os.path.exists(LIGHT_STATUS)

    if log_exists:
        # read from LIGHT_STATUS
        light_status = open(LIGHT_STATUS, "r").readlines()[0] == "True\n" # pylint: disable=[W1514, R1732]
    else:
        print("LIGHT_STATUS not present")
        light_status = get_smartthings_device_status()
        # write to LIGHT_STATUS
        write_light_status(light_status)

    return light_status
