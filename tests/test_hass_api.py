from config import CONFIG, PYTEST
from hass.api import get_entity_state, toggle_entity_state

# from mock_data import MOCK_ENTITY_STATE


def test_get_entity_state(mocker):
    # Create a mock response object with a .json() method that returns the mock data
    mock_response = mocker.MagicMock()
    mock_response.json.return_value = PYTEST["MOCK_ENTITY_STATE"]

    # Patch 'requests.get' to return the mock response
    mocker.patch("requests.get", return_value=mock_response)

    # Call the function
    entity_state = get_entity_state(CONFIG["HASS_ENTITY_ID"])

    # assertions
    assert isinstance(entity_state, dict)
    assert entity_state["state"] in ["on", "off"]
    assert "entity_id" in entity_state.keys()
    assert "state" in entity_state.keys()
    assert "last_changed" in entity_state.keys()


def test_toggle_entity_state(mocker):
    # Create a mock response object with a .json() method that returns the mock data
    mock_response = mocker.MagicMock()
    mock_response.json.return_value = PYTEST["MOCK_ENTITY_STATE"]

    # Patch 'requests.get' to return the mock response
    mocker.patch("requests.get", return_value=mock_response)

    # Call the function
    toggled_state = toggle_entity_state(CONFIG["HASS_ENTITY_ID"])

    # assertions
    assert isinstance(toggled_state, dict)
    assert toggled_state["state"] in ["on", "off"]
