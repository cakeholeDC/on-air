from host.actions import is_trigger_app_running, get_and_log_device_status
from api.actions import turn_off_and_log_status, turn_on_and_log_status

if __name__ == '__main__':
    WEBCAM_STATUS = is_trigger_app_running()
    print(f"isAppOpen = {WEBCAM_STATUS}")

    lightStatus = get_and_log_device_status()
    print(f"isLightOn = {lightStatus}")

    if WEBCAM_STATUS and lightStatus:
        # do nothing
        pass
    if not WEBCAM_STATUS and not lightStatus:
        # do nothing
        pass
    elif WEBCAM_STATUS and not lightStatus:
        # turn light on
        turn_on_and_log_status()
    elif not WEBCAM_STATUS and lightStatus:
        # turn light off
        turn_off_and_log_status()
    else:
        # something has gone terribly wrong
        pass
