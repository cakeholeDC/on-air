import pytest

from homekit.control import homekit_device_off, homekit_device_on, homekit_device_state


@pytest.mark.skip(reason="test hangs...")
def test_homekit_device_state():
    device_state = homekit_device_state()
    assert isinstance(device_state, bool)
    assert device_state is False


@pytest.mark.skip(reason="test hangs...")
def test_homekit_device_on():
    homekit_device_on()
    # pass
    device_state = homekit_device_state()
    assert isinstance(device_state, bool)
    assert device_state is True


@pytest.mark.skip(reason="test hangs...")
def test_homekit_device_off():
    homekit_device_off()
    # pass
    device_state = homekit_device_state()
    assert isinstance(device_state, bool)
    assert device_state is False
