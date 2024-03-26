#!/bin/bash

cd "$(dirname "$0")" || exit;

PATH="$PATH:$HOME/.local/bin"

poetry run python3 off_air.py
