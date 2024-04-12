from pathlib import Path

from host.cache import read_cache, update_cache

GIT_ROOT = Path(__file__).resolve().parent.parent

PYTEST_CACHE = GIT_ROOT / ".pytest.cache"


def test_update_cache():
    response = update_cache(PYTEST_CACHE, True)
    assert response[0] == PYTEST_CACHE
    assert response[1] is True


def test_read_cache():
    update_cache(PYTEST_CACHE, False)
    response = read_cache(PYTEST_CACHE)
    assert response[0] == PYTEST_CACHE
    assert response[1] is False


def teardown_module():
    PYTEST_CACHE.unlink(missing_ok=True)
