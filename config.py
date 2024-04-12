import json
import os
from pathlib import Path

from dotenv import load_dotenv

load_dotenv()

GIT_ROOT = Path(__file__).resolve().parent

CONFIG = {
    "ENABLE_VIDEO": os.getenv("ENABLE_VIDEO", "False") == "True",
    "ENABLE_AUDIO": os.getenv("ENABLE_AUDIO", "False") == "True",
    "TRIGGER_APPS": json.loads(os.getenv("TRIGGER_APPS", "[]")),
    "DEVICE_CACHE": str(GIT_ROOT / os.getenv("DEVICE_CACHE", ".cache.device")),
    "VIDEO_CACHE": str(GIT_ROOT / os.getenv("VIDEO_CACHE", ".cache.video")),
    "SHORTCUT_ON": os.getenv("SHORTCUT_ON"),
    "SHORTCUT_OFF": os.getenv("SHORTCUT_OFF"),
    "SHORTCUT_STATE": os.getenv("SHORTCUT_STATE"),
}
