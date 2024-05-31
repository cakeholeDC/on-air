import platform
import subprocess

import psutil

from config import CONFIG, GIT_ROOT
from host.cache import read_cache, update_cache
from host.logger import logger

VIDEO_STATE_SH = "macos_camera_state.sh"
AUDIO_STATE_SH = "macos_audio_engine_state.sh"


def is_video_active() -> bool:
    """
    Parses the last 5 minutes of system logs for webcam power state entries:
    - USB webcam = VDCAssistant_Power_State
    - Builtin Webcam = AppleH13CamIn::setGetPowerStateGated

    If no log entry is found, checks the video_cache for camera status.

    Returns a boolean; writes the boolean to the video_cache.
    """
    logger.debug("👀 Checking macos camera state")
    if platform.system() != "Darwin":
        logger.debug("Platform is not Darwin. Returning False")
        return False

    cmd_result = subprocess.run(
        [str(GIT_ROOT / "scripts" / VIDEO_STATE_SH)],
        stdout=subprocess.PIPE,
        check=False,
    )
    video_state = cmd_result.stdout.decode().strip()

    # handle for no log activity in the last n minutes
    if video_state == "check_cache":
        cache_file, video_cache = read_cache(  # pylint: disable=W0612
            CONFIG["VIDEO_CACHE"]
        )
        video_state = str(video_cache)

    logger.debug(f"🎥 {video_state=}")  # pylint: disable=W1203
    video_state_bool = video_state == "True"
    update_cache(CONFIG["VIDEO_CACHE"], video_state_bool)

    return video_state_bool


def is_audio_active() -> bool:
    """
    Parses the io registry (ioreg) for 'IOAudioEngineState'

    Returns a boolean; does not cache.
    """
    logger.debug("👀 Checking macos audio engine state")
    if platform.system() != "Darwin":
        logger.debug("Platform is not Darwin. Returning False")
        return False

    cmd_result = subprocess.run(
        [str(GIT_ROOT / "scripts" / AUDIO_STATE_SH)],
        stdout=subprocess.PIPE,
        check=False,
    )
    audio_state = cmd_result.stdout.decode().strip()
    logger.debug(f"🎙️ {audio_state=}")  # pylint: disable=W1203

    return audio_state == "True"


def is_app_open() -> bool:
    """
    Iterates running processes using psutil. Checks if any of the TRIGGER_APPS are currently running.

    Returns a boolean.
    """
    # ? TODO: should we cache this?
    logger.debug("👀 Checking host process list")
    if platform.system() != "Darwin":
        logger.debug("Platform is not Darwin. Returning False")
        return False

    found_processes = []
    # TODO: can this be made faster by iterating the trigger apps and checking if they're running
    for proc in psutil.process_iter():
        try:
            process_name = proc.name()
            if process_name in CONFIG["TRIGGER_APPS"]:
                found_processes.append(process_name)
                logger.debug(  # pylint: disable=W1203
                    f"🏃‍➡️ Running process named '{process_name}' found!"
                )
        except (psutil.NoSuchProcess, psutil.AccessDenied, psutil.ZombieProcess):
            pass

    if len(found_processes) > 0:  # pylint: disable=R1705
        logger.debug("💿 app_state='True'")
        return True
    else:
        logger.debug("🙈 No running processes found")
        logger.debug("💿 app_state='False'")
        return False
