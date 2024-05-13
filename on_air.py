from config import CONFIG
from host.cache import update_cache
from host.colors import COLOR_NONE, GREEN, RED
from host.control import onair
from host.logger import logger


def run():
    logger.info("💻 Manually invoked 'onair'")
    onair()
    update_cache(CONFIG["DEVICE_CACHE"], True)

    print(f"{GREEN}------------------")
    print(f"{GREEN}| --- {RED}ON AIR {GREEN}--- |")
    print(f"{GREEN}------------------{COLOR_NONE}")


if __name__ == "__main__":
    run()
