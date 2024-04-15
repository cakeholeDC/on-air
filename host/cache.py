from pathlib import Path

from host.logger import logger


def update_cache(cache: str, value: bool) -> tuple[Path, bool]:
    """
    Writes a boolean to a specified cache file (Path)

    Returns (Path, bool)
    """
    cache_path = Path(cache)
    # pylint: disable-next=W1203
    logger.debug(f"🤑 Updating cache: {cache_path} => {value}")
    with open(cache_path, "w+", encoding="utf-8") as cache_file:
        cache_file.write(str(value))
        cache_file.close()
    return (cache_path, str(value) == "True")


def read_cache(cache: Path) -> tuple[Path, bool]:
    """
    Reads a boolean from a specified cache file (Path)

    Returns (Path, bool)
    """
    cache_path = Path(cache)
    with open(cache_path, "r", encoding="utf-8") as cache_file:
        value = cache_file.readlines()[0]
    # pylint: disable-next=W1203
    logger.debug(f"👀 Reading cache: {cache_path} = {value}")
    return (cache_path, value == "True")
