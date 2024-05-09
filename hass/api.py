import json
import requests

from config import CONFIG

from host.logger import logger

LOG_MODULE="🏠 HASS:"

HASS_BASE_API=f"{CONFIG["HASS_SERVER_URL"]}/api"

def get_entity_state(entity_id: str) -> dict:
    url = f"{HASS_BASE_API}/states/{entity_id}"

    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {CONFIG["HASS_API_TOKEN"]}'
    }
    payload = {}

    response = requests.request(
        "GET",
        url,
        headers=headers,
        data=payload,
    )

    entity_state = json.loads(response.text)
    # TODO: logger.info
    print(f"{LOG_MODULE} {entity_id} => GET => {entity_state['state']}")
    return entity_state

# TODO: consider 'service' name due to api url.
def toggle_entity_state(entity_id: str) -> dict:
    url = f"{HASS_BASE_API}/services/homeassistant/toggle"

    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {CONFIG["HASS_API_TOKEN"]}'
    }
    payload = json.dumps({
        "entity_id": entity_id
    })

    response = requests.request(
        "POST",
        url,
        headers=headers,
        data=payload,
    )

    toggled_state = json.loads(response.text)[0]
    # TODO: logger.info
    print(f"{LOG_MODULE} {entity_id} => SET => {toggled_state['state']}")
    return toggled_state
