import subprocess

from config import CONFIG
from host.logger import logger


def homekit_device_on() -> None:
    """
    Runs an Apple shortcut to turn on the Homekit Device
    """
    logger.debug("🏡 homekit device on")
    # Run the apple shortcut to turn the device on
    output = subprocess.run(
        args=["/usr/bin/shortcuts", "run", CONFIG["SHORTCUT_ON"]],
        capture_output=True,
        check=False,
    )
    return output


def homekit_device_off() -> None:
    """
    Runs an Apple shortcut to turn off the Homekit Device
    """
    logger.debug("🏡 homekit device off")
    # Run the apple shortcut to turn the device off
    output = subprocess.run(
        args=["/usr/bin/shortcuts", "run", CONFIG["SHORTCUT_OFF"]],
        capture_output=True,
        check=False,
    )
    return output


def homekit_device_state() -> bool:
    """
    Runs an Apple shortcut to get the current state of the Homekit Device
    """
    # Run the apple shortcut to get the device's state
    output = subprocess.run(
        args=["/usr/bin/shortcuts", "run", CONFIG["SHORTCUT_STATE"]],
        capture_output=True,
        check=False,
    )
    status = output.stdout.decode()
    logger.debug(f"🏡 homekit device {status=}")  # pylint: disable=W1203
    return status == "Yes"
