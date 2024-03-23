# from api.actions import turn_off_and_log_status
from homekit.control import homekit_off
from api.colors import BLACK, DARK_GRAY, COLOR_NONE

if __name__ == "__main__":
    # TODO: move homekit into the log function
    homekit_off()
    # turn_off_and_log_status()
    print(f"{DARK_GRAY}------------------")
    print(f"{DARK_GRAY}| --- {BLACK}ON AIR{DARK_GRAY} --- |")
    print(f"{DARK_GRAY}------------------{COLOR_NONE}")
