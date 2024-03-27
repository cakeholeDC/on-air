# from homekit.control import homekit_device_off
import os

from host.actions import turn_off_and_log_status
from host.colors import BLACK, COLOR_NONE, DARK_GRAY

if __name__ == "__main__":
    # homekit_device_off()
    os.system("shortcuts run 'offair'")
    turn_off_and_log_status()
    print(f"{DARK_GRAY}------------------")
    print(f"{DARK_GRAY}| --- {BLACK}ON AIR{DARK_GRAY} --- |")
    print(f"{DARK_GRAY}------------------{COLOR_NONE}")
