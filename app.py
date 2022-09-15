from host.actions import get_and_log_vdc_status, is_trigger_app_running, get_and_log_device_status, get_vdc_assistant_power_state
from api.actions import turn_off_and_log_status, turn_on_and_log_status

if __name__ == '__main__':
    ######## IMPORTANT NOTE: THIS MAY ONLY WORK WITH EXTERNAL WEBCAMS. DIG DEEPER INTO THE ACTUAL SCRIPT.
    
    # WEBCAM_STATUS = get_vdc_assistant_power_state() ## SHOULD RETURN BOOLEAN
    USE_WEBCAM = True
    if USE_WEBCAM:
        # ### turn_on_and_log_status()
        WEBCAM_STATUS = get_and_log_vdc_status()
        
        if WEBCAM_STATUS:
            turn_on_and_log_status()
        else: 
            turn_off_and_log_status()
    else:

        # todo: work in vdc state
        APP_STATUS = is_trigger_app_running()
        print(f"isAppOpen = {APP_STATUS}")

        lightStatus = get_and_log_device_status()
        print(f"isLightOn = {lightStatus}")

        if APP_STATUS and lightStatus:
            # do nothing
            pass
        if not APP_STATUS and not lightStatus:
            # do nothing
            pass
        elif APP_STATUS and not lightStatus:
            # turn light on
            turn_on_and_log_status()
        elif not APP_STATUS and lightStatus:
            # turn light off
            turn_off_and_log_status()
        else:
            # something has gone terribly wrong
            pass
