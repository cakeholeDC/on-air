import os
from pathlib import Path

SHORTCUT_ON = os.getenv('SHORTCUT_ON')
SHORTCUT_OFF = os.getenv('SHORTCUT_OFF')
SHORTCUT_STATE = os.getenv('SHORTCUT_STATE')

# TODO: run shortcut.
def homekit_device_on():
    # Run the apple shortcut to turn the device on
    os.system(f"shortcuts run '{SHORTCUT_ON}'")

def homekit_device_off():
    # Run the apple shortcut to turn the device off
    os.system(f"shortcuts run '{SHORTCUT_OFF}'")

def homekit_device_state():
    # Run the apple shortcut to get the device's state
    status = os.system(f"shortcuts run '{SHORTCUT_STATE}'")
    print(f"homekit_device_state = {(status == 1)}")
    return (status == 1)

if __name__ == "__main__":
    homekit_device_state()