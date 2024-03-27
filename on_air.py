import os

# from homekit.control import homekit_device_on
from host.actions import turn_on_and_log_status
from host.colors import COLOR_NONE, GREEN, RED

if __name__ == "__main__":
    # homekit_device_on()
    os.system("shortcuts run 'onair'")
    turn_on_and_log_status()
    print(f"{GREEN}------------------")
    print(f"{GREEN}| --- {RED}ON AIR {GREEN}--- |")
    print(f"{GREEN}------------------{COLOR_NONE}")
    # os.system("printenv")
