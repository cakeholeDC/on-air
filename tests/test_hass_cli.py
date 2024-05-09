import pytest

from hass.cli import toggle_entity, get_state

# TODO: config
device = "switch.on_air"

def test_hass_toggle_entity():
    toggle_entity(device)
    # assert isinstance(device_state, bool)
    # assert device_state is False

def test_get_state():
    get_state(device)
