from config import CONFIG
from homekit.control import homekit_device_off
from host.cache import update_cache
from host.colors import BLACK, COLOR_NONE, DARK_GRAY

if __name__ == "__main__":
    homekit_device_off()
    update_cache(CONFIG["DEVICE_CACHE"], False)

    print(f"{DARK_GRAY}------------------")
    print(f"{DARK_GRAY}| --- {BLACK}ON AIR{DARK_GRAY} --- |")
    print(f"{DARK_GRAY}------------------{COLOR_NONE}")
