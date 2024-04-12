import logging
from logging.handlers import RotatingFileHandler

from config import GIT_ROOT

# TODO: configurable LOG_LEVEL

# Inspired by: https://stackoverflow.com/a/24505345
log_formatter = logging.Formatter(
    "%(asctime)s [%(levelname)s] %(filename)s::%(funcName)s(%(lineno)d) %(message)s"
)

logFile = GIT_ROOT / "onair.log"

my_handler = RotatingFileHandler(
    logFile, mode="a", maxBytes=10000, backupCount=2, encoding="utf-8", delay=False
)
my_handler.setFormatter(log_formatter)
my_handler.setLevel(logging.DEBUG)

logger = logging.getLogger("onair")
logger.setLevel(logging.DEBUG)

logger.addHandler(my_handler)
