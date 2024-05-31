import os

import pytest

from config import CONFIG
from hass.api import (
    get_entity_state,
    toggle_entity_state,
    turn_off_entity,
    turn_on_entity,
)
from hass.fixtures import MOCK_API


@pytest.mark.skipif(
    os.environ.get("MOCK_API") == "True",
    reason="Skip when running in CI",
)
def test_get_entity_state():
    entity_state = get_entity_state(CONFIG["HASS_ENTITY_ID"])
    # assertions
    assert isinstance(entity_state, dict)
    assert entity_state.keys() == MOCK_API["STATES_ENTITY_RESPONSE"].keys()


@pytest.mark.skipif(
    os.environ.get("MOCK_API") != "True",
    reason="Skip when running locally",
)
def test_get_entity_state_mocked(mocker):
    # Create a mock response object with a .json() method that returns the mock data
    mock_response = mocker.MagicMock()
    mock_response.json.return_value = MOCK_API["STATES_ENTITY_RESPONSE"]

    # Patch 'requests.get' to return the mock response
    mocker.patch("requests.get", return_value=mock_response)

    # Call the function
    entity_state = get_entity_state(CONFIG["HASS_ENTITY_ID"])

    # assertions
    assert isinstance(entity_state, dict)
    assert entity_state == MOCK_API["STATES_ENTITY_RESPONSE"]


@pytest.mark.skipif(
    os.environ.get("MOCK_API") == "True",
    reason="Skip when running in CI",
)
def test_toggle_entity_state():
    toggled_state = toggle_entity_state(CONFIG["HASS_ENTITY_ID"])
    # assertions
    assert isinstance(toggled_state[0], dict)
    assert toggled_state[0]["state"] in ["on", "off"]


def test_get_entity_state_with_ci_mock(mocker):
    if os.environ.get("MOCK_API") == "True":
        # Create a mock response object with a .json() method that returns the mock data
        mock_response = mocker.MagicMock()
        mock_response.json.return_value = MOCK_API["STATES_ENTITY_RESPONSE"]

        # Patch 'requests.get' to return the mock response
        mocker.patch("requests.get", return_value=mock_response)

    # Call the function
    entity_state = get_entity_state(CONFIG["HASS_ENTITY_ID"])

    # assertions
    assert isinstance(entity_state, dict)
    assert entity_state.keys() == MOCK_API["STATES_ENTITY_RESPONSE"].keys()
    if os.environ.get("MOCK_API") == "True":
        assert entity_state == MOCK_API["STATES_ENTITY_RESPONSE"]


def test_toggle_entity_state_with_ci_mock(mocker):
    if os.environ.get("MOCK_API") == "True":
        # Create a mock response object with a .json() method that returns the mock data
        mock_response = mocker.MagicMock()
        mock_response.json.return_value = MOCK_API["SERVICES_TOGGLE_RESPONSE"]

        # Patch 'requests.get' to return the mock response
        mocker.patch("requests.post", return_value=mock_response)

    # Call the function
    toggled_state = toggle_entity_state(CONFIG["HASS_ENTITY_ID"])

    # assertions
    assert isinstance(toggled_state[0], dict)
    assert toggled_state[0].keys() == MOCK_API["SERVICES_TOGGLE_RESPONSE"][0].keys()
    if os.environ.get("MOCK_API") == "True":
        assert toggled_state == MOCK_API["SERVICES_TOGGLE_RESPONSE"]


def test_turn_on_entity_with_ci_mock(mocker):
    if os.environ.get("MOCK_API") == "True":
        # Create a mock response object with a .json() method that returns the mock data
        mock_response = mocker.MagicMock()
        mock_response.json.return_value = MOCK_API["SERVICES_ENTITY_ON_RESPONSE"]

        # Patch 'requests.get' to return the mock response
        mocker.patch("requests.post", return_value=mock_response)

    # Call the function
    response = turn_on_entity(CONFIG["HASS_ENTITY_ID"])

    # assertions
    assert isinstance(response[0], dict)
    assert response[0].keys() == MOCK_API["SERVICES_ENTITY_ON_RESPONSE"][0].keys()
    assert response[0]["state"] == "on"
    if os.environ.get("MOCK_API") == "True":
        assert response == MOCK_API["SERVICES_ENTITY_ON_RESPONSE"]


def test_turn_off_entity_with_ci_mock(mocker):
    if os.environ.get("MOCK_API") == "True":
        # Create a mock response object with a .json() method that returns the mock data
        mock_response = mocker.MagicMock()
        mock_response.json.return_value = MOCK_API["SERVICES_ENTITY_OFF_RESPONSE"]

        # Patch 'requests.get' to return the mock response
        mocker.patch("requests.post", return_value=mock_response)

    # Call the function
    response = turn_off_entity(CONFIG["HASS_ENTITY_ID"])

    # assertions
    assert isinstance(response[0], dict)
    assert response[0].keys() == MOCK_API["SERVICES_ENTITY_OFF_RESPONSE"][0].keys()
    assert response[0]["state"] == "off"
    if os.environ.get("MOCK_API") == "True":
        assert response == MOCK_API["SERVICES_ENTITY_OFF_RESPONSE"]
