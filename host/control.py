import sys

from config import CONFIG
from hass.api import turn_off_entity, turn_on_entity
from homekit.api import homekit_device_off, homekit_device_on
from host.logger import logger


def onair():
    """
    Runs the ON command for the configured INTEGRATION
    """
    if CONFIG["INTEGRATION"] == "homekit":
        homekit_device_on()
    elif CONFIG["INTEGRATION"] == "hass":
        turn_on_entity(CONFIG["HASS_ENTITY_ID"])
    elif CONFIG["INTEGRATION"] == "smartthings":
        logger.error("❌ Smartthings not supported at this time")
        sys.exit(1)
    logger.info("🌕 Turned ON Device")


def offair():
    """
    Runs the OFF command for the configured INTEGRATION
    """
    if CONFIG["INTEGRATION"] == "homekit":
        homekit_device_off()
    elif CONFIG["INTEGRATION"] == "hass":
        turn_off_entity(CONFIG["HASS_ENTITY_ID"])
    elif CONFIG["INTEGRATION"] == "smartthings":
        logger.error("❌ Smartthings not supported at this time")
        sys.exit(1)
    logger.info("🌑 Turned OFF Device")
