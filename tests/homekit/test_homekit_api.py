from homekit.api import homekit_device_off, homekit_device_on, homekit_device_state


def test_homekit_device_state(mocker):
    # Create a mock response object that returns the mock data
    mock_response = mocker.MagicMock()
    mock_response.return_value = "No"

    # Patch 'subprocess.run' to return the mock response
    mocker.patch("subprocess.run", return_value=mock_response)

    device_state = homekit_device_state()
    assert isinstance(device_state, bool)
    assert device_state is False


def test_homekit_device_on(mocker):
    # Create a mock response object that returns the mock data
    mock_response = mocker.MagicMock()
    mock_response.return_value = True

    # Patch 'subprocess.run' to return the mock response
    mocker.patch("subprocess.run", return_value=mock_response)

    # pylint: disable-next=E1101
    device_state = homekit_device_on().return_value
    assert isinstance(device_state, bool)
    assert device_state is True


def test_homekit_device_off(mocker):
    # Create a mock response object that returns the mock data
    mock_response = mocker.MagicMock()
    mock_response.return_value = False

    # Patch 'subprocess.run' to return the mock response
    mocker.patch("subprocess.run", return_value=mock_response)

    # pylint: disable-next=E1101
    device_state = homekit_device_off().return_value
    assert isinstance(device_state, bool)
    assert device_state is False
