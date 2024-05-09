import subprocess
from config import CONFIG
from host.logger import logger

# TODO: move ID into config.
def toggle_entity(entity_id) -> None:
    # TODO: use logger statements.
    print(CONFIG)
    # TODO: Config secrets. IP, Token, etc.
    print("🎚 toggling hass entity")
    # TODO: hass-cli is not official. https://github.com/home-assistant-ecosystem/home-assistant-cli
        # but it is supported => https://www.home-assistant.io/blog/2019/02/04/introducing-home-assistant-cli/
    # TODO: Try this one: https://github.com/home-assistant/cli
        # => i think this is a part of the HASS OS specifically. 
    subprocess.run(
        args=[
            "hass-cli",
            "--server",
            "http://10.88.216.136:8123",
            "--token",
            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjOTI3ZjA2ODdmYjc0MjJmODY1MTliZGUyMmM5ZWEzNCIsImlhdCI6MTcxNTE5MDk5MSwiZXhwIjoyMDMwNTUwOTkxfQ.EMAmNj0ykrnL22Tyu-71fOY3PGl6FhcmZqiCaP1hoTU",
            "state",
            "toggle",
            entity_id,
        ],
        check=False
    )

def get_state(entity_id):
    """
    $ hass-cli state get switch.desk_lamp

    ENTITY            DESCRIPTION    STATE    CHANGED
    switch.desk_lamp  Desk Lamp      on       2024-05-09T17:04:46.615088+00:00
    """
    result = subprocess.run(
        args=[
            "hass-cli",
            "--server",
            "http://10.88.216.136:8123",
            "--token",
            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjOTI3ZjA2ODdmYjc0MjJmODY1MTliZGUyMmM5ZWEzNCIsImlhdCI6MTcxNTE5MDk5MSwiZXhwIjoyMDMwNTUwOTkxfQ.EMAmNj0ykrnL22Tyu-71fOY3PGl6FhcmZqiCaP1hoTU",
            "state",
            "get",
            entity_id,
            # "|",
            # "grep",
            # entity_id
        ],
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE, 
        check=False
    )
    print(result.stdout)