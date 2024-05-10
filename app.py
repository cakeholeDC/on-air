import sys

from config import CONFIG, SUPPORTED_INTEGRATIONS
from host.cache import read_cache, update_cache
from host.control import offair, onair
from host.logger import logger
from host.scraping import is_app_open, is_audio_active, is_video_active


def handle_logic(
    video_state: bool = False,
    audio_state: bool = False,
    app_state: bool = False,
    device_state: bool = False,
):
    """
    Handles host state, and determines if the device state should change.
    """
    # light should be on, and is on
    if (video_state or audio_state or app_state) and device_state:
        logger.info("👍 Light should be ON, and is ON")
        # do nothing

    # light should be on, but is off
    elif (video_state or audio_state or app_state) and not device_state:
        logger.info("👎 Light should be ON, but is OFF")
        # turn light on
        logger.info("🌕 Turning ON Device")
        onair()
        # update cache
        update_cache(CONFIG["DEVICE_CACHE"], True)

    # light should be off, but is on
    elif (not video_state and not audio_state and not app_state) and device_state:
        logger.info("👎 Light should be OFF, but is ON")
        # turn light off
        logger.info("🌑 Turning OFF Device")
        offair()
        # update cache
        update_cache(CONFIG["DEVICE_CACHE"], False)

    # light should be off, and is off
    elif (not video_state and not audio_state and not app_state) and not device_state:
        logger.info("👍 Light should be OFF, and is OFF")
        # do nothing

    # something has gone terribly wrong
    else:
        logger.error("something has gone terribly wrong")
        # do nothing


def validate_config(config: dict):
    logger.debug(f"🛠️ CONFIG={config}")  # pylint: disable=W1203
    if config["INTEGRATION"] not in SUPPORTED_INTEGRATIONS:
        # pylint: disable-next=W1203
        logger.error(
            f"❌ Unsupported integration: {config['INTEGRATION']} - must be one of {SUPPORTED_INTEGRATIONS}"
        )
        sys.exit(1)
    else:
        # pylint: disable-next=W1203
        logger.info(f"🤖🏠 Using Integration: {config['INTEGRATION']}")


def run():
    """
    Runs the application.

    Parses the config and gets the current host state.

    Passes host state to the device logic handler.
    """
    logger.info("🎙️🚨 Running onair")
    validate_config(CONFIG)

    # Get video IO state if config.video is enabled
    video_state = is_video_active() if CONFIG["ENABLE_VIDEO"] else False

    # Get audio IO state if config.audio is enabled
    audio_state = is_audio_active() if CONFIG["ENABLE_AUDIO"] else False

    # Are any config.trigger_apps running?
    app_state = is_app_open()

    # Get the power state of the device
    # TODO: integrate 'homekit.control::homekit_device_state' and '.cache.device'
    cache_path, device_state = read_cache(  # pylint: disable=W0612
        CONFIG["DEVICE_CACHE"]
    )

    # determine whether the light should be on or off.
    handle_logic(
        video_state=video_state,
        audio_state=audio_state,
        app_state=app_state,
        device_state=device_state,
    )


if __name__ == "__main__":
    run()
