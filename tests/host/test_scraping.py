from host.scraping import is_app_open, is_audio_active, is_video_active


def test_is_audio_active():
    audio_state = is_audio_active()
    assert audio_state is False
    assert isinstance(audio_state, bool)


def test_is_video_active():
    video_state = is_video_active()
    assert video_state is False
    assert isinstance(video_state, bool)


def test_is_app_open():
    app_state = is_app_open()
    assert app_state is False
    assert isinstance(app_state, bool)
