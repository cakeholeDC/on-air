import json

import requests

from config import CONFIG

# from host.logger import logger

LOG_MODULE = "🏠 HASS:"

HASS_BASE_API = f"{CONFIG['HASS_SERVER_URL']}/api"


def get_entity_state(entity_id: str) -> dict:
    url = f"{HASS_BASE_API}/states/{entity_id}"

    headers = {
        "Content-Type": "application/json",
        "Authorization": f'Bearer {CONFIG["HASS_API_TOKEN"]}',
    }
    payload = {}

    response = requests.get(
        url,
        headers=headers,
        data=payload,
        timeout=5,
    )

    return response.json()


# TODO: consider 'service' name due to api url.
def toggle_entity_state(entity_id: str) -> dict:
    url = f"{HASS_BASE_API}/services/homeassistant/toggle"

    headers = {
        "Content-Type": "application/json",
        "Authorization": f'Bearer {CONFIG["HASS_API_TOKEN"]}',
    }
    payload = json.dumps({"entity_id": entity_id})

    response = requests.post(
        url,
        headers=headers,
        data=payload,
        timeout=5,
    )

    return response.json()


def turn_on_entity(entity_id: str) -> dict:
    url = f"{HASS_BASE_API}/services/homeassistant/turn_on"

    headers = {
        "Content-Type": "application/json",
        "Authorization": f'Bearer {CONFIG["HASS_API_TOKEN"]}',
    }
    payload = json.dumps({"entity_id": entity_id})

    response = requests.post(
        url,
        headers=headers,
        data=payload,
        timeout=5,
    )

    return response.json()


def turn_off_entity(entity_id: str) -> dict:
    url = f"{HASS_BASE_API}/services/homeassistant/turn_off"

    headers = {
        "Content-Type": "application/json",
        "Authorization": f'Bearer {CONFIG["HASS_API_TOKEN"]}',
    }
    payload = json.dumps({"entity_id": entity_id})

    response = requests.post(
        url,
        headers=headers,
        data=payload,
        timeout=5,
    )

    return response.json()
