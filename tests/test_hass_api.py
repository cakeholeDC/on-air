import os

import pytest

from config import CONFIG, PYTEST
from hass.api import get_entity_state, toggle_entity_state

# from mock_data import MOCK_ENTITY_STATE


@pytest.mark.skipif(
    os.environ.get("GITHUB_MOCK_API") == "True", reason="Skip when running in CI"
)
def test_get_entity_state():
    entity_state = get_entity_state(CONFIG["HASS_ENTITY_ID"])

    # assertions
    assert isinstance(entity_state, dict)
    assert entity_state.keys() == PYTEST["MOCK_ENTITY_STATE"].keys()


@pytest.mark.skipif(
    os.environ.get("GITHUB_MOCK_API") != "True", reason="Skip when running locally"
)
def test_get_entity_state_mocked(mocker):
    # Create a mock response object with a .json() method that returns the mock data
    mock_response = mocker.MagicMock()
    mock_response.json.return_value = PYTEST["MOCK_ENTITY_STATE"]

    # Patch 'requests.get' to return the mock response
    mocker.patch("requests.get", return_value=mock_response)

    # Call the function
    entity_state = get_entity_state(CONFIG["HASS_ENTITY_ID"])

    # assertions
    assert isinstance(entity_state, dict)
    assert entity_state == PYTEST["MOCK_ENTITY_STATE"]


# TODO: mock this function in CI
@pytest.mark.skipif(
    os.environ.get("GITHUB_MOCK_API") == "True", reason="Skip when running in CI"
)
def test_toggle_entity_state():
    # Create a mock response object with a .json() method that returns the mock data
    # mock_response = mocker.MagicMock()
    # mock_response.json.return_value = PYTEST["MOCK_ENTITY_STATE"]

    # # Patch 'requests.get' to return the mock response
    # mocker.patch("requests.get", return_value=mock_response)
    # print(mock_response)
    # Call the function
    toggled_state = toggle_entity_state(CONFIG["HASS_ENTITY_ID"])
    print(toggled_state)
    # assertions
    assert isinstance(toggled_state[0], dict)
    assert toggled_state[0]["state"] in ["on", "off"]
