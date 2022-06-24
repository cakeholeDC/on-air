#!/bin/bash
cd "$(dirname "$0")" || exit;

PATH="$HOME/.poetry/bin:$HOME/.pyenv/shims:$PATH"

poetry run python3 on_air.py
