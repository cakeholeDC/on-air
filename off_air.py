from api.actions import turn_off_and_log_status
from api.colors import BLACK, DARK_GRAY, COLOR_NONE

if __name__ == "__main__":
    turn_off_and_log_status()
    print(f"{DARK_GRAY}------------------")
    print(f"{DARK_GRAY}| --- {BLACK}ON AIR{DARK_GRAY} --- |")
    print(f"{DARK_GRAY}------------------{COLOR_NONE}")
