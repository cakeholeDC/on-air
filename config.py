import json
import os
from pathlib import Path

from dotenv import load_dotenv

load_dotenv()

GIT_ROOT = Path(__file__).resolve().parent

SUPPORTED_INTEGRATIONS = ["homekit", "hass", "smartthings"]

CONFIG = {
    # App Configuration
    "ENABLE_VIDEO": os.getenv("ENABLE_VIDEO", "False") == "True",
    "ENABLE_AUDIO": os.getenv("ENABLE_AUDIO", "False") == "True",
    "TRIGGER_APPS": json.loads(os.getenv("TRIGGER_APPS", "[]")),
    # Must be one of ["homekit", "hass", "smartthings"]
    "INTEGRATION": os.getenv("INTEGRATION", "homekit"),
    # Host Cache
    "DEVICE_CACHE": str(GIT_ROOT / os.getenv("DEVICE_CACHE", ".cache.device")),
    "VIDEO_CACHE": str(GIT_ROOT / os.getenv("VIDEO_CACHE", ".cache.video")),
    #  Home Assistant
    "HASS_SERVER_URL": os.getenv("HASS_SERVER_URL", None),
    "HASS_API_TOKEN": os.getenv("HASS_API_TOKEN", None),
    "HASS_ENTITY_ID": os.getenv("HASS_ENTITY_ID", None),
    # Apple HomeKit
    "SHORTCUT_ON": os.getenv("SHORTCUT_ON", None),
    "SHORTCUT_OFF": os.getenv("SHORTCUT_OFF", None),
    "SHORTCUT_STATE": os.getenv("SHORTCUT_STATE", None),
}
