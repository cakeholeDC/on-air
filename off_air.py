from config import CONFIG
from host.cache import update_cache
from host.colors import BLACK, COLOR_NONE, DARK_GRAY
from host.control import offair
from host.logger import logger


def run():
    logger.info("💻 Manually invoked: 'offair'")
    offair()
    update_cache(CONFIG["DEVICE_CACHE"], False)

    print(f"{DARK_GRAY}------------------")
    print(f"{DARK_GRAY}| --- {BLACK}ON AIR{DARK_GRAY} --- |")
    print(f"{DARK_GRAY}------------------{COLOR_NONE}")


if __name__ == "__main__":
    run()
