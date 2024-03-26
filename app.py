import os

from host.actions import (
  get_and_log_device_status,
  get_and_log_vdc_status,
  is_trigger_app_running,
  turn_off_and_log_status,
  turn_on_and_log_status
)

USE_WEBCAM = (os.getenv('USE_WEBCAM', "False") == "True")

if __name__ == '__main__':
    if USE_WEBCAM:
        ## NOTE: This may only work with external usb webcams. Do more testing and research.
        WEBCAM_STATUS = get_and_log_vdc_status()
        print(f"isWebcamEnabled = {WEBCAM_STATUS}")
    else:
        WEBCAM_STATUS = False

    # check if any of the trigger apps are running.
    APP_STATUS = is_trigger_app_running()
    print(f"isAppOpen = {APP_STATUS}")

    # get the current status of the indicator light.
    DEVICE_STATUS = get_and_log_device_status()
    print(f"isLightOn = {DEVICE_STATUS}")

    # light should be on, and is on
    if (WEBCAM_STATUS or APP_STATUS) and DEVICE_STATUS:
        # do nothing
        pass

    # light should be on, but is off
    elif (WEBCAM_STATUS or APP_STATUS) and not DEVICE_STATUS:
        # turn light on
        turn_on_and_log_status()

    # light should be off, but is on
    elif (not WEBCAM_STATUS and not APP_STATUS) and DEVICE_STATUS:
        # turn light off
        turn_off_and_log_status()

    # light should be off, and is off
    elif (not WEBCAM_STATUS and not APP_STATUS) and not DEVICE_STATUS:
        # do nothing
        pass

    else:
        # something has gone terribly wrong
        # do nothing
        pass
