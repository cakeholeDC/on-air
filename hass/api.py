import json

import requests

from config import CONFIG

HASS_BASE_API = f"{CONFIG['HASS_SERVER_URL']}/api"


def get_entity_state(entity_id: str) -> dict:
    """
    GET the state/entity_id object

    Returns an entity/state object
    """
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


def toggle_entity_state(entity_id: str) -> dict:
    """
    POST a services/toggle request for a entity_id(s)

    Returns a list of entity/state objects.
    """
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
    """
    POST a services/turn_on request for a entity_id(s)

    Returns a list of entity/state objects.
    """
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
    """
    POST a services/turn_off request for a entity_id(s)

    Returns a list of entity/state objects.
    """
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
