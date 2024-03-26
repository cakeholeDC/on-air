import os
from pathlib import Path

from dotenv import load_dotenv

load_dotenv()

GIT_ROOT = Path(__file__).resolve().parent.parent
LIGHT_STATUS = os.environ.get("LIGHT_STATUS")
CAMERA_STATUS = os.getenv("CAMERA_STATUS")


def write_light_status(status):
    """
    Writes `True` or `False` to the local log file
    """
    light_state = open(
        f"{GIT_ROOT / LIGHT_STATUS}", "w"
    )  # pylint: disable=[W1514, R1732]
    light_state.write(str(status))
    light_state.write("\n")
    light_state.close()


def write_camera_status(status):
    """
    Writes `True` or `False` to the local log file
    """
    camera_state = open(GIT_ROOT / CAMERA_STATUS, "w")  # pylint: disable=[W1514, R1732]
    camera_state.write(str(status))
    camera_state.write("\n")
    camera_state.close()


if __name__ == "__main__":
    write_light_status(True)
