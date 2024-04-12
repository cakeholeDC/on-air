from config import CONFIG
from homekit.control import homekit_device_on
from host.cache import update_cache
from host.colors import COLOR_NONE, GREEN, RED

if __name__ == "__main__":
    homekit_device_on()
    update_cache(CONFIG["DEVICE_CACHE"], True)

    print(f"{GREEN}------------------")
    print(f"{GREEN}| --- {RED}ON AIR {GREEN}--- |")
    print(f"{GREEN}------------------{COLOR_NONE}")
