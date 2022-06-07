from host.actions import is_trigger_app_running, get_and_log_device_status
from api.actions import turn_off_and_log_status, turn_on_and_log_status

if __name__ == '__main__':
    webcamStatus = is_trigger_app_running()
    print(f"isAppOpen = {webcamStatus}")

    lightStatus = get_and_log_device_status()
    print(f"isLightOn = {lightStatus}")

    if webcamStatus and lightStatus:
        # do nothing
        pass
    if not webcamStatus and not lightStatus:
        # do nothing
        pass
    elif webcamStatus and not lightStatus:
        # turn light on
        turn_on_and_log_status()
    elif not webcamStatus and lightStatus:
        # turn light off
        turn_off_and_log_status()
    else: 
        # something has gone terribly wrong
        pass