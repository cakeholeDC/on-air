import asyncio
from api.LightControl import ( 
    lights_on, 
    lights_off, 
    get_device_status,
    list_devices,
    lights_object, 
    device_name
)
from logger import write_log_file

loop = asyncio.get_event_loop()

def turn_on_and_log_status():
    '''
    turns on the light and writes the status to the log file.
    '''
    loop.run_until_complete( lights_on(lights_object, device_name) )
    write_log_file(True)

def turn_off_and_log_status():
    '''
    turns off the light and writes the status to the log file.
    '''
    loop.run_until_complete( lights_off(lights_object, device_name) )
    write_log_file(False)

def get_smartthings_device_list():
    return loop.run_until_complete( list_devices() )

def get_smartthings_device_status(device_name):
    return loop.run_until_complete( get_device_status(device_name) )