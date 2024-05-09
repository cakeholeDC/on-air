from config import CONFIG
from hass.api import toggle_entity_state, get_entity_state

def test_get_entity_state():
    entity_state = get_entity_state(CONFIG["HASS_ENTITY_ID"])
    assert isinstance(entity_state, dict)
    assert entity_state['state'] in ["on","off"]
    assert "entity_id" in entity_state.keys()
    assert "state" in entity_state.keys()
    assert "last_changed" in entity_state.keys()

def test_toggle_entity_state():
    toggled_state = toggle_entity_state(CONFIG["HASS_ENTITY_ID"])
    assert isinstance(toggled_state, dict)
    assert toggled_state['state'] in ["on","off"]